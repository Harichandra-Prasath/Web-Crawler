package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Harichandra-Prasath/Web-Crawler/utils"
	"golang.org/x/net/html"
)

var (
	failed  int
	success int
)

func worker(q *utils.Queue, wg *sync.WaitGroup, curr string, root string) {
	defer wg.Done()
	res, err := http.Get(curr)
	if err != nil {
		failed += 1
		return
	}
	success += 1
	doc, _ := html.Parse(res.Body)
	utils.Getallnodes(doc, wg, q, root)

}

func main() {
	q := utils.GetQueue()
	root := "https://transform.tools/"
	q.Append(root)
	wg := new(sync.WaitGroup)
	start_time := time.Now()
	for len(q.Elements) > 0 {
		//fmt.Println(q.Elements)
		for i := 0; i < len(q.Elements); i++ {
			wg.Add(1)
			curr := q.Popleft()
			//fmt.Println("Popped element: ", curr)
			go worker(q, wg, curr, root)

		}
		wg.Wait()
	}
	fmt.Println(len(q.Elements))
	fmt.Println(len(q.History))
	fmt.Println(failed)
	fmt.Println(success)
	fmt.Println(time.Since(start_time))
}
