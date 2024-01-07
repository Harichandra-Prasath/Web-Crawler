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

func Getallnodes(node *html.Node, q *Queue) {
	var crawl func(*html.Node)
	crawl = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					_, ok := q.History[attr.Val]
					if strings.HasPrefix(attr.Val, "https://scrapeme.live/shop/") && !ok {
						q.Append(attr.Val)
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
