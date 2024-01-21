package utils

import (
	"fmt"
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
	q.mu.RLock()
	if _, ok := q.History[element]; !ok {
		q.mu.RUnlock()
		q.mu.Lock()
		if _, ok := q.History[element]; !ok {
			q.History[element] = true
			q.Elements = append(q.Elements, element)
		}
		q.mu.Unlock()
	} else {
		q.mu.RUnlock()
	}
}

func (q *Queue) Popleft() string {
	element := q.Elements[0]
	q.Elements = q.Elements[1:]
	return element

}

func GetQueue() *Queue {
	q := &Queue{}
	q.History = make(map[string]bool)
	return q
}

func Getallnodes(node *html.Node, q *Queue, root string, same_domian bool, curr string) {
	var url string
	var crawl func(*html.Node)
	crawl = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					if same_domian {
						if strings.HasPrefix(attr.Val, "/") { //to capture the realtive paths
							url = root + strings.TrimPrefix(attr.Val, "/")
							q.Append(url)
						} else if strings.HasPrefix(attr.Val, root) {
							url = attr.Val
							q.Append(url)
						}

					} else {
						if strings.HasPrefix(attr.Val, "/") {
							if strings.HasPrefix(curr, root) {
								url = root + strings.TrimPrefix(attr.Val, "/")
								q.Append(url)
							} else {
								url = curr + strings.TrimPrefix(attr.Val, "/")
								q.Append(url)
							}
						} else {
							if !strings.HasPrefix(attr.Val, "#") && attr.Val != "" { //to cut off sections
								fmt.Println(curr, attr.Val)
								q.Append(attr.Val)
							}
						}

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
