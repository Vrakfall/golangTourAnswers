package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)[:]
	wordsCount := make(map[string]int)

	for _, word := range words {
		wordsCount[word]++
	}

	return wordsCount
}

func main() {
	wc.Test(WordCount)
}
