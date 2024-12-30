package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) < 3 {
		fmt.Println("not enough arguments provided\n Usage: webcrawler <website> <max_concurrency> <max_pages>")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided\n Usage: webcrawler <website> <max_concurrency> <max_pages>")
		os.Exit(1)
	}

	base_url := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("error parsing max concurrency")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("error parsing max pages")
		os.Exit(1)
	}

	cfg, err := configure(base_url, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println("Error configuring crawler:\n", err)
		os.Exit(1)
	}

	fmt.Println("starting crawl of:", base_url)

	cfg.wg.Add(1)
	go cfg.crawlPage(base_url)
	cfg.wg.Wait()

	fmt.Println("\nCrawling Results: ")
	for page, count := range cfg.pages {
		fmt.Printf(" - %d: %s\n", count, page)
	}
}
