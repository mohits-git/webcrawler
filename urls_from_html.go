package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error parsing base URL: %v", err)
	}

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("Error parsing HTML: %v", err)
	}

	var urls []string
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n == nil {
			return
		}

		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					u, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("Error parsing URL: %v\n", err)
						continue
					}
					resolvedUrl := baseUrl.ResolveReference(u)
					urls = append(urls, resolvedUrl.String())
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	for i, u := range urls {
		if strings.HasPrefix(u, "/") {
			urls[i] = rawBaseURL + u
		}
	}

	return urls, nil
}
