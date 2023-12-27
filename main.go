package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func Getallnodes(node *html.Node, q *Queue, target string) {
	var crawl func(*html.Node)
	crawl = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == target {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					if strings.HasPrefix(attr.Val, "https://scrapeme.live/shop/") && !q.contains(attr.Val) {
						q.append(attr.Val)
						fmt.Print(attr.Val)
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
	n := 1
	for n <= 1 {
		fmt.Print(n)
		curr := q.popleft()
		res, err := http.Get(curr)
		if err != nil {
			fmt.Print("error occured")
		}
		doc, _ := html.Parse(res.Body)

		Getallnodes(doc, q, "a")
		n = n + 1
	}
}

type Queue struct {
	Elements []string
}

func (q *Queue) append(element string) {
	q.Elements = append(q.Elements, element)
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
	for _, el := range q.Elements {
		if el == element {
			return true
		}
	}
	return false
}
