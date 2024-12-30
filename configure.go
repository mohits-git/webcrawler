package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseUrl            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
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

func configure(base_url string, maxConcurrency int) (*config, error) {
	baseUrl, err := url.Parse(base_url)
	if err != nil {
		return nil, fmt.Errorf("Error parsing base URL:\n %v", err)
	}

	cfg := config{
		pages:              make(map[string]int),
		baseUrl:            baseUrl,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	return &cfg, nil
}
