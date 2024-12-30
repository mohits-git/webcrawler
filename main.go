package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
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

	baseUrl, err := url.Parse(base_url)
	if err != nil {
		fmt.Println("Error parsing base URL:\n", err)
		os.Exit(1)
	}
	
	maxConcurrency := 10

	cfg := config{
		pages:						make(map[string]int),
		baseUrl:					baseUrl,
		mu:							&sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:							&sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(base_url)
	cfg.wg.Wait()

	fmt.Println("\nCrawling Results: ")
	fmt.Println("Internal Pages Links: ")
	for page, count := range cfg.pages {
		fmt.Printf(" - %d: %s\n", count, page)
	}
}
