package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputUrl string) (string, error) {
    parsedUrl, err := url.Parse(inputUrl)
    if err != nil {
        return "", fmt.Errorf("Error parsing URL: %w", err)
    }

    normalizedUrl := parsedUrl.Host + parsedUrl.Path
    normalizedUrl = strings.ToLower(normalizedUrl)
    normalizedUrl = strings.TrimSuffix(normalizedUrl, "/")
    return normalizedUrl, nil
}
