package main

import (
	"fmt"
	"strings"
)

func countWords(s string) map[string]int {
	wordCountMap := make(map[string]int)
	for _, word := range strings.Fields(s) {
		wordCountMap[word]++
	}
	return wordCountMap
}

// Do not change the code in the main function
func main() {
	s := "go is fun and go is fast"
	fmt.Println(countWords(s))
	t := "this is a test this is only a test"
	fmt.Println(countWords(t))
}
