package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func Getallnodes(node *html.Node, q *Queue) {
	var crawl func(*html.Node)
	crawl = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					if strings.HasPrefix(attr.Val, "https://scrapeme.live/shop/") && !q.contains(attr.Val) {
						q.append(attr.Val)
					}

				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawl(child)
		}
	}

	crawl(node)
}

var (
	failed int
)

func worker(q *Queue, wg *sync.WaitGroup) {
	defer wg.Done()
	curr := q.popleft()
	if curr == "" {
		return
	}
	res, err := http.Get(curr)
	if err != nil {
		failed += 1
		return
	}
	doc, _ := html.Parse(res.Body)
	Getallnodes(doc, q)
}

func main() {
	q := &Queue{}
	q.append("https://scrapeme.live/shop/")
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
	fmt.Println(len(q.history))
	fmt.Println(failed)
	fmt.Println(time.Since(start))

}

type Queue struct {
	Elements []string
	history  []string
}

func (q *Queue) append(element string) {
	q.Elements = append(q.Elements, element)
	q.history = append(q.history, element)
}

func (q *Queue) popleft() string {
	if len(q.Elements) > 0 {
		element := q.Elements[0]
		if len(q.Elements) == 1 {
			q.Elements = nil
			return element
		}
		q.Elements = q.Elements[1:]
		return element
	} else {
		return ""
	}

}

func (q *Queue) contains(element string) bool {
	for _, el := range q.history {
		if el == element {
			return true
		}
	}
	return false
}
