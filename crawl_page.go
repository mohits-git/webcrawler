package main

import (
	"fmt"
	"net/url"
	"strings"
)

func crawlPage(rawBaseUrl, rawCurrentUrl string, pages map[string]int) {
	pasrsedBaseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		fmt.Printf("Error parsing base URL: %v\n", err)
		return
	}

	parsedCurrentUrl, err := url.Parse(rawCurrentUrl)
	if err != nil {
		fmt.Printf("Error parsing current URL: %v\n", err)
		return
	}

	if parsedCurrentUrl.Hostname() != pasrsedBaseUrl.Hostname() {
		return
	}

  currentUrl, err := normalizeURL(rawCurrentUrl)
  if err != nil {
		fmt.Println("Error normalizing URL:", rawCurrentUrl, "\n", err)
    return
  }

  if !strings.HasPrefix(rawCurrentUrl, rawBaseUrl) {
    return
  }

  if c, ok := pages[currentUrl]; ok {
    pages[currentUrl] = c+1;
    return;
  } else {
    pages[currentUrl] = 1
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
    crawlPage(rawBaseUrl, link, pages)
  }

  fmt.Println("Crawl complete for", rawCurrentUrl)
}
