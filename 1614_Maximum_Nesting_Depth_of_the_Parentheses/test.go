package main

import (
	"fmt"
)

func main() {
	testdata := "((()))"
	result := maxDepth(testdata)
	fmt.Printf("result is : %d", result)
}

// 贪心 前一天与后一天有点麻烦
func maxDepth(s string) int {
	cemetery := 0
	result := 0
	for _, i := range s {
		if i == '(' {
			result++
			if result > cemetery {
				cemetery = result
			}
		} else if i == ')' {
			result--
		}
	}
	return cemetery
}
