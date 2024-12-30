# Web Crawler

This is a simple web crawler CLI tool that crawls a website and prints a report of all the internal links on that website. Built with Go.

## Getting Started

**Clone Repository**

```bash
git clone https://github.com/mohits-git/webcrawler
cd webcrawler
```

**Install CLI Tool**

```bash
go install
```

or

**Build and Run CLI Tool**

```bash
go build
./webcrawler <url> <max_concurrency> <max_pages>
```

## Usage

```bash
webcrawler <url> <max_concurrency> <max_pages>
```

- `url` (string): The URL of the website to crawl. (e.g. https://example.com)
- `max_concurrency` (int): The maximum number of concurrent requests to make. (e.g. 10)
- `max_pages` (int): The maximum number of pages to crawl or to nest into. (e.g. 100)

## Example

```bash
webcrawler https://news.ycombinator.com 3 25
```

output:

```
.
.
...crawling logs
.
.
=============================
REPORT for https://news.ycombinator.com
=============================
Found 1 internal links to news.ycombinator.com/asknew
Found 1 internal links to news.ycombinator.com/formatdoc
Found 1 internal links to news.ycombinator.com/newpoll
Found 1 internal links to news.ycombinator.com/showhn.html
Found 2 internal links to news.ycombinator.com/shownew
Found 5 internal links to news.ycombinator.com/front
Found 5 internal links to news.ycombinator.com/jobs
Found 5 internal links to news.ycombinator.com/lists
Found 5 internal links to news.ycombinator.com/login
Found 5 internal links to news.ycombinator.com/news
Found 5 internal links to news.ycombinator.com/security.html
Found 5 internal links to news.ycombinator.com/submit
Found 6 internal links to news.ycombinator.com/newcomments
Found 6 internal links to news.ycombinator.com/newsfaq.html
Found 7 internal links to news.ycombinator.com
Found 7 internal links to news.ycombinator.com/ask
Found 7 internal links to news.ycombinator.com/newest
Found 7 internal links to news.ycombinator.com/newsguidelines.html
Found 7 internal links to news.ycombinator.com/show
Found 30 internal links to news.ycombinator.com/context
Found 60 internal links to news.ycombinator.com/hide
Found 87 internal links to news.ycombinator.com/from
Found 149 internal links to news.ycombinator.com/user
Found 149 internal links to news.ycombinator.com/vote
Found 372 internal links to news.ycombinator.com/item
```

## Demo 

https://github.com/user-attachments/assets/d3dd8dd4-4ce5-4b54-a5f5-2991ab914283

