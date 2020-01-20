package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  var count = map[string]int{}
  for _, word := range strings.Fields(s) {
    if _, found := count[word]; found {
      count[word] += 1
    } else {
      count[word] = 1
    }
  }
  return count
}

func main() {
  wc.Test(WordCount)
}
