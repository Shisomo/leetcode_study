package main

import (
	"fmt"
)

func main() {
	data := "ababababa"
	result := removePalindromeSub(data)
	fmt.Printf("result is : %d", result)
}

func removePalindromeSub(s string) int {
	n := len(s) - 1
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i] {
			return 2
		}
	}
	return 1
}
