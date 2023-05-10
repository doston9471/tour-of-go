package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	// Run the test suite
	wc.Test(WordCount)

	// Additional example usage
	text := "I am learning Go programming language and Go is fun!"
	result := WordCount(text)

	// Print the word count
	for word, count := range result {
		fmt.Printf("%s: %d\n", word, count)
	}
}
