package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "this is a apple"
	data2 := "this is a banana"
	result := uncommonFromSentences(data, data2)
	fmt.Printf("result is : %v", result)
}

func uncommonFromSentences(s1 string, s2 string) []string {
	cemetery := make(map[string]int)
	initMap := func(list string) {
		words := strings.Split(list, " ")
		for _, word := range words {
			cemetery[word]++
		}
	}
	initMap(s1)
	initMap(s2)
	var result []string
	for word, n := range cemetery {
		if n == 1 {
			result = append(result, word)
		}
	}
	return result
}
