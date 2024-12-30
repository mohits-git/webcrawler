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
		fmt.Printf("Error getting HTML from %s :\n %v\n", rawCurrentUrl, err)
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

	fmt.Println("Crawl complete for", rawCurrentUrl)
}

func (cfg *config) addPageVisits(rawCurrentUrl string) (isFirst bool) {
	currentUrl, err := normalizeURL(rawCurrentUrl)
	if err != nil {
		fmt.Println("Error normalizing URL:", rawCurrentUrl, "\n", err)
		return
	}

	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if c, ok := cfg.pages[currentUrl]; ok {
		cfg.pages[currentUrl] = c + 1
		return
	}

	cfg.pages[currentUrl] = 1
	return true
}
