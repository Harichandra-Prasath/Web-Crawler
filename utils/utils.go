package utils

import (
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Queue struct {
	Elements []string
	History  map[string]bool
	mu       sync.RWMutex
}

func (q *Queue) Append(element string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Elements = append(q.Elements, element)
	q.History[element] = true
}

func (q *Queue) Check(element string) bool {
	q.mu.RLock()
	defer q.mu.RUnlock()
	_, ok := q.History[element]
	return ok

}

func (q *Queue) Popleft() string {

	if len(q.Elements) == 0 {
		return ""
	} else {
		element := q.Elements[0]
		q.Elements = q.Elements[1:]
		return element
	}

}

func GetQueue() *Queue {
	q := &Queue{}
	q.History = make(map[string]bool)
	return q
}

func Getallnodes(node *html.Node, wg *sync.WaitGroup, q *Queue, root string) {
	var url string
	var crawl func(*html.Node)
	crawl = func(node *html.Node) {

		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {

					if strings.HasPrefix(attr.Val, "/") { //to capture the realtive paths
						url = root + strings.TrimPrefix(attr.Val, "/")
					} else {
						url = attr.Val
					}
					ok := q.Check(url)
					if strings.HasPrefix(url, root) && !ok {
						q.Append(url)
						//fmt.Println(url)
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
