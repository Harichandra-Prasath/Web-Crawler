package main

import (
	"os"

	"github.com/Harichandra-Prasath/Web-Crawler/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
