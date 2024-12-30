package main

import (
	"fmt"
	"os"
)

const maxConcurrency = 10

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

	cfg, err := configure(base_url, maxConcurrency)
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
