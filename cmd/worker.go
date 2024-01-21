package cmd

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

func worker(q *utils.Queue, wg *sync.WaitGroup, curr string, root string, same_domain bool) {
	defer wg.Done()
	res, err := http.Get(curr)
	if err != nil {
		failed += 1
		return
	}
	success += 1
	doc, _ := html.Parse(res.Body)
	utils.Getallnodes(doc, q, root, same_domain)

}

func Crawl(root string, same_domain bool, generate bool, depth int) {
	q := utils.GetQueue()
	links := bytes.Buffer{}
	if generate {
		file, _ := os.Create("links.txt")
		defer end(file, &links)
	}

	q.Append(root)
	wg := new(sync.WaitGroup)
	start_time := time.Now()
	for i := 0; i < depth; i++ {
		if len(q.Elements) == 0 {
			break
		}
		links.WriteString(fmt.Sprintf("Level %d:\n", i+1))
		length := len(q.Elements)
		for i := 0; i < length; i++ {
			wg.Add(1)
			curr := q.Popleft()
			go worker(q, wg, curr, root, same_domain)
			links.WriteString(fmt.Sprintf("            - %s\n", curr))
		}
		wg.Wait()

	}
	fmt.Println("Current Links in the Queue: ", len(q.Elements))
	fmt.Println("Total Links Scraped:        ", len(q.History))
	fmt.Println("Failures:                   ", failed)
	fmt.Println("No of Links Crawled:        ", success)
	fmt.Println("Total time:                 ", time.Since(start_time))

}

func end(file *os.File, links *bytes.Buffer) {
	file.Write(links.Bytes())
	file.Close()
}
