package main

import (
	"fmt"
	"slices"
	"strings"
)

func printReport(pages map[string]int, baseUrl string) {
	fmt.Println("=============================")
	fmt.Println("REPORT for", baseUrl)
	fmt.Println("=============================")

	type freqUrl struct {
		freq int
		url  string
	}
	var report []freqUrl

	for k, v := range pages {
		report = append(report, freqUrl{v, k})
	}

	slices.SortStableFunc(report, func(a, b freqUrl) int {
		return strings.Compare(a.url, b.url)
	})

	slices.SortStableFunc(report, func(a, b freqUrl) int {
		return a.freq - b.freq
	})

  for _, v := range report {
    fmt.Printf("Found %d internal links to %s\n", v.freq, v.url)
  }
}
