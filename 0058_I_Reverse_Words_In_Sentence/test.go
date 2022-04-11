package main

import (
	"fmt"
	"strings"
)

func main() {
	testdata := "hello world!"
	result := reverseWords(testdata)
	fmt.Printf("result is : %d", result)
}

func reverseWords(s string) string {
	if s == "" || strings.Count(s, " ") == len(s) {
		return ""
	}
	i, j := len(s)-1, len(s)-1
	var result []byte
	for i >= 0 {
		if (s[i] == ' ') && s[j] != ' ' {
			result = append(result, s[i+1:j+1]...)
			result = append(result, ' ')
			j = i
		} else if i == 0 && s[j] != ' ' {
			result = append(result, s[i:j+1]...)
			result = append(result, ' ')
		}
		if s[j] == ' ' {
			j--
		}
		i--
	}
	return string(result[:len(result)-1])
}
