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

func worker(q *utils.Queue, wg *sync.WaitGroup, curr string) {
	defer wg.Done()
	res, err := http.Get(curr)
	if err != nil {
		failed += 1
		return
	}
	success += 1
	doc, _ := html.Parse(res.Body)
	utils.Getallnodes(doc, q)
}

func main() {
	q := utils.GetQueue()
	q.Append("https://scrapeme.live/shop/")
	wg := new(sync.WaitGroup)
	start_time := time.Now()
	for len(q.Elements) > 0 {
		for i := 0; i < len(q.Elements); i++ {
			wg.Add(1)
			curr := q.Popleft()
			go worker(q, wg, curr)
		}
		wg.Wait()
	}
	fmt.Println(len(q.Elements))
	fmt.Println(len(q.History))
	fmt.Println(failed)
	fmt.Println(success)
	fmt.Println(time.Since(start_time))
}
