package main

import (
	"fmt"
	"os"
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
  fmt.Println("starting crawl of:", base_url)

	pageHtml, err := getHTML(base_url)
	if err != nil {
		fmt.Println("Error getting HTML:\n", err)
		os.Exit(1)
	}

	fmt.Println(pageHtml)

	// links, err := getURLsFromHTML(pageHtml, base_url)
	// if err != nil {
	// 	fmt.Println("Error getting URLs from HTML:\n", err)
	// 	os.Exit(1)
	// }
	//
	// fmt.Println("Links found: ")
	// for _, link := range links {
	// 	fmt.Println(" - ", link)
	// }
	//
	// fmt.Println("Crawl complete")
}
