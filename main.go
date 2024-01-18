package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
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
	utils.Getallnodes(doc, q, root)

}

func main() {
	q := utils.GetQueue()
	file, _ := os.Create("links.txt")
	defer file.Close()
	links := bytes.Buffer{}
	root := "https://transform.tools/"
	q.Append(root)
	wg := new(sync.WaitGroup)
	start_time := time.Now()
	for i := 0; i < 3; i++ {
		links.WriteString(fmt.Sprintf("Level %d:\n", i))
		length := len(q.Elements)
		for i := 0; i < length; i++ {
			wg.Add(1)
			curr := q.Popleft()
			go worker(q, wg, curr, root)
			links.WriteString(fmt.Sprintf("            - %s\n", curr))
		}
		wg.Wait()

	}
	fmt.Println(len(q.Elements))
	fmt.Println(len(q.History))
	fmt.Println(failed)
	fmt.Println(success)
	fmt.Println(time.Since(start_time))
	file.Write(links.Bytes())

}
