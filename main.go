package main

import (
	"fmt"
	"os"
)

func main() {
  args := os.Args[1:]
  if len(args) == 0 {
    fmt.Println("no website provided")
    os.Exit(1)
  }

  if len(args) > 1 {
    fmt.Println("too many arguments provided")
    os.Exit(1)
  }

  base_url := args[0]
  fmt.Println("starting crawl of:", base_url)
}
