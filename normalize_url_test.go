package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{
			name:     "Remove Scheme",
			inputURL: "http://example.com/path/",
			expected: "example.com/path",
		},
		{
			name:     "Remove Trailing Slash",
			inputURL: "http://example.com/path/",
			expected: "example.com/path",
		},
		{
			name:     "Lowercase Capital Letters",
			inputURL: "http://EXAMPLE.com/path/",
			expected: "example.com/path",
		},
		{
			name:     "Remove Scheme and Trailing Slash and Lowercase Capital Letters",
			inputURL: "http://EXAMPLE.com/path/",
			expected: "example.com/path",
		},
		{
			name:          "Handle Invalid URL",
			inputURL:      "http://[::1",
			expected:      "",
			errorContains: "Error parsing URL",
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := normalizeURL(tt.inputURL)
			if err != nil && !strings.Contains(err.Error(), tt.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tt.name, err)
				return
			} else if err != nil && tt.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tt.name, err)
				return
			} else if err == nil && tt.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%s', got nil", i, tt.name, tt.errorContains)
				return
			}

			if actual != tt.expected {
				t.Errorf("Test %v - '%s' FAIL: expected '%s', got '%s'", i, tt.name, tt.expected, actual)
			}
		})
	}

}
