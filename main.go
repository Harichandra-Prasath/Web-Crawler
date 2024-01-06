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
	failed int
)

func worker(q *utils.Queue, wg *sync.WaitGroup) {
	defer wg.Done()
	curr := q.Popleft()
	if curr == "" {
		return
	}
	res, err := http.Get(curr)
	if err != nil {
		failed += 1
		return
	}
	doc, _ := html.Parse(res.Body)
	utils.Getallnodes(doc, q)
}

func main() {
	q := utils.GetQueue()
	q.Append("https://scrapeme.live/shop/")
	wg := new(sync.WaitGroup)
	start := time.Now()
	for len(q.Elements) > 0 {
		for i := 0; i < len(q.Elements); i++ {
			wg.Add(1)
			go worker(q, wg)
		}
		wg.Wait()
	}
	fmt.Println(len(q.Elements))
	fmt.Println(len(q.History))
	fmt.Println(failed)
	fmt.Println(time.Since(start))

}
