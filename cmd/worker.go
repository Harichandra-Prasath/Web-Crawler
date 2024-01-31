package cmd

import (
	"bytes"
	"fmt"
	"log/slog"
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
	utils.Getallnodes(doc, q, root, same_domain, curr)

}

func Intiate(root string, root_relative bool, generate bool, depth int) {
	start_time := time.Now()
	q := utils.GetQueue()
	q.Append(root)
	fmt.Println()
	slog.Info("Crawling initiated for root", "url", root)
	if generate {
		home, _ := os.UserHomeDir()
		file, _ := os.Create(fmt.Sprintf("%s/output.txt", home))
		links := bytes.Buffer{}
		deepen_gen(q, depth, root, root_relative, &links)
		defer end(file, &links, home)
	} else {
		deepen(q, depth, root, root_relative)
	}
	fmt.Println()
	fmt.Println("Current Links in the Queue   : ", len(q.Elements))
	fmt.Println("Total Links Scraped          : ", len(q.History))
	fmt.Println("Failures                     : ", failed)
	fmt.Println("No of Links Crawled          : ", success)
	fmt.Println("Total time                   : ", time.Since(start_time))
	fmt.Println()

}

func end(file *os.File, links *bytes.Buffer, home string) {
	slog.Info("Generating the file...")
	_, err := file.Write(links.Bytes())
	if err != nil {
		slog.Info("Error in generating the file")
		slog.Error(err.Error())
	} else {
		slog.Info("File Successfuly Generated")
		slog.Info("output.txt generated at", "location", home)
		fmt.Println()
	}

	file.Close()
}

func deepen_gen(q *utils.Queue, depth int, root string, same_domain bool, links *bytes.Buffer) {
	wg := new(sync.WaitGroup)

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
		slog.Info("Crawling Done for ", "level", i+1)
	}
}

func deepen(q *utils.Queue, depth int, root string, same_domain bool) {
	wg := new(sync.WaitGroup)

	for i := 0; i < depth; i++ {
		if len(q.Elements) == 0 {
			break
		}
		length := len(q.Elements)
		for i := 0; i < length; i++ {
			wg.Add(1)
			curr := q.Popleft()
			go worker(q, wg, curr, root, same_domain)
		}
		wg.Wait()

	}
}
