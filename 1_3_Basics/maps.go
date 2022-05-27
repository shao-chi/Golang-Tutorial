package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	word_list := strings.Fields(s)
	word_count := make(map[string]int)

	for _, str := range word_list {
		word_count[str] += 1
	}
	return word_count
}

func main() {
	wc.Test(WordCount)
}
