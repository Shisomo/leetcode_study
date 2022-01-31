package main

import (
	"fmt"
)

func main() {
	data := []string{
		"hello",
		"world",
		"df"}
	result := shortestDistance(data, "hello", "world")
	fmt.Println(result)
}

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	word1Index, word2Index := -1, -1
	result := len(wordsDict)
	for i, n := range wordsDict {
		if n == word1 {
			word1Index = i
		}
		if n == word2 {
			word2Index = i
		}
		if word1Index != -1 && word2Index != -1 {
			if word1Index-word2Index > 0 {
				if word1Index-word2Index < result {
					result = word1Index - word2Index
				}
			} else {
				if word2Index-word1Index < result {
					result = word2Index - word1Index
				}
			}
		}
	}
	return result
}
