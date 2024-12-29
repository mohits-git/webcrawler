package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputBody     string
		inputURL      string
		expected      []string
		errorContains string
	}{
		{
			name:      "relative URLs",
			inputURL:  "https://blog.example.com",
			inputBody: "<html><body><a href='/path/one'>Example.com</a></body></html>",
			expected:  []string{"https://blog.example.com/path/one"},
		},
		{
			name:      "absolute URLs",
			inputURL:  "https://blog.example.com",
			inputBody: "<html><body><a href='https://other.com/path/one'>Example.com</a></body></html>",
			expected:  []string{"https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.example.com",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Example.com</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Example.com</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.example.com/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no URLs",
			inputURL: "https://blog.example.com",
			inputBody: `
<html>
	<body>
		<span>Example.com</span>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "Other HTML tags with hrefs",
			inputURL: "https://blog.example.com",
			inputBody: `
<html>
	<body>
		<anchor href="/path/one">
			<span>Example.com</span>
		</anchor>
		<anchor href="https://other.com/path/one">
			<span>Example.com</span>
		</anchor>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "invalid HTML",
			inputURL: "https://blog.example.com",
			inputBody: `
<html body>
	<a href="/path/one">
		<span>Example.com</span>
	</a>
</html body>
`,
			expected: []string{"https://blog.example.com/path/one"},
		},
		{
			name:          "invalid base URL",
			inputURL:      "://blog.example.com",
			inputBody:     "<html><body><a href='/path/one'>Example.com</a></body></html>",
			errorContains: "Error parsing base URL",
		},
		{
			name:     "HTML with invalid href attribute",
			inputURL: "https://blog.example.com",
			inputBody: `
<html>
	<body>
		<a href="://other">
			<span>Example.com</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actual, err := getURLsFromHTML(test.inputBody, test.inputURL)
			if err != nil {
				if test.errorContains == "" {
					tt.Fatalf("\nunexpected error: %v", err)
				}
				if !strings.Contains(err.Error(), test.errorContains) {
					tt.Fatalf("\nexpected error to contain %q, got %q", test.errorContains, err.Error())
				}
				return
			}
			if test.errorContains != "" {
				tt.Fatalf("\nexpected error to contain %q, got nil, \n and urls: %v", test.errorContains, actual)
				return
			}
			if len(actual) != len(test.expected) {
				tt.Fatalf("\nexpected %d urls, got %d\n urls: %v", len(test.expected), len(actual), actual)
				return
			}
			if !reflect.DeepEqual(actual, test.expected) {
				tt.Fatalf("\nexpected %v, got %v", test.expected, actual)
				return
			}
		})
	}
}
