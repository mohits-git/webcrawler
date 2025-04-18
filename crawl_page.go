package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentUrl string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()

	pagesVisited := cfg.getPagesLength()
	if pagesVisited >= cfg.maxPages {
		return
	}

	parsedCurrentUrl, err := url.Parse(rawCurrentUrl)
	if err != nil {
		fmt.Printf("Error parsing current URL: %v\n", err)
		return
	}

	if parsedCurrentUrl.Hostname() != cfg.baseUrl.Hostname() {
		return
	}

	if isFirst := cfg.addPageVisits(rawCurrentUrl); !isFirst {
		return
	}

	fmt.Println("Crawling", rawCurrentUrl)

	pageHtml, err := getHTML(rawCurrentUrl)
	if err != nil {
		fmt.Printf("html not found for %s\n", rawCurrentUrl)
		return
	}

	links, err := getURLsFromHTML(pageHtml, rawCurrentUrl)
	if err != nil {
		fmt.Printf("Error getting URLs from HTML of %s :\n %v\n", rawCurrentUrl, err)
		return
	}

	for _, link := range links {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}
}
