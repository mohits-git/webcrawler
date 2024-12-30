package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func getHTML(rawUrl string) (string, error){
  parsedUrl, err := url.Parse(rawUrl)
  if err != nil {
    return "", fmt.Errorf("failed to parse url: %v", err)
  }

  client := &http.Client{
    Timeout: 5 * time.Second,
  }

  req, err := http.NewRequest("GET", parsedUrl.String(), nil)
  if err != nil {
    return  "", fmt.Errorf("failed to create request: %v", err)
  }

  resp, err := client.Do(req)
  if err != nil {
    return "", fmt.Errorf("failed to get response: %v", err)
  }
  defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("response status code was %d", resp.StatusCode)
	}

  if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "text/html") {
    return "", fmt.Errorf("response was not an HTML page (Content-Type: %s)", ct)
  }

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return "", fmt.Errorf("failed to read response body: %v", err)
  }

  return string(body), nil
}
