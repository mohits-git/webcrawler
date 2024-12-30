package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
  args := os.Args[1:]
  if len(args) == 0 {
    fmt.Println("no website provided")
    os.Exit(1)
  }

  if len(args) > 1 {
    fmt.Println("too many arguments provided")
    os.Exit(1)
  }

  base_url := args[0]
	base_url = strings.TrimSuffix(base_url, "/")

  fmt.Println("starting crawl of:", base_url)

	pages := make(map[string]int)
	crawlPage(base_url, base_url, pages)

	fmt.Println("\nCrawling Results: ")
	fmt.Println("Internal Pages Links: ")
	for page, count := range pages {
		fmt.Printf(" - %d: %s\n", count, page)
	}
}
