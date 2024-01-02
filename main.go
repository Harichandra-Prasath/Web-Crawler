package main

import (
	"fmt"
	"net/http"
	"strings"

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
						fmt.Printf("Added element %s\n", attr.Val)
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

func main() {
	q := &Queue{}
	q.append("https://scrapeme.live/shop/")
	for len(q.Elements) > 0 {
		curr := q.popleft()
		fmt.Printf("Popped element %s\n", curr)
		res, err := http.Get(curr)
		if err != nil {
			fmt.Print("error occured")
		}
		doc, _ := html.Parse(res.Body)

		Getallnodes(doc, q)

	}
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
	element := q.Elements[0]
	if len(q.Elements) == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element
}

func (q *Queue) contains(element string) bool {
	for _, el := range q.history {
		if el == element {
			return true
		}
	}
	return false
}
