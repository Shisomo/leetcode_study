package main

import "fmt"

func main() {
	s := "abcdefg"
	//fmt.Println(s[0])
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	//var a [256]int
	//fmt.Printf("%d",a)
}
func lengthOfLongestSubstring(s string) int {
	result, left, right := 0, 0, 0
	var cemetery [256]int
	for right < len(s) {
		if right < len(s) && cemetery[s[right]] == 0 {
			cemetery[s[right]]++
			right++
		} else {
			cemetery[s[left]]--
			left++
		}
		result = max(result, right-left)

	}
	return result
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
