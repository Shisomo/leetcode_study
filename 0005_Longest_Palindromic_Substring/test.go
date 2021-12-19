package main

import "fmt"

func main() {
	testString := "abcdefga"
	result := longestPalindrome(testString)
	fmt.Printf("result is : %f", result)
}
func longestPalindrome(s string) string {
	var cemetery [256]int
	left, right := 0, 0
	for right < len(s) {
		if cemetery[s[right]] == 0 {
			right++
			cemetery[s[right]]++
		} else {
			left++
			return s
		}
	}
	return s
}
