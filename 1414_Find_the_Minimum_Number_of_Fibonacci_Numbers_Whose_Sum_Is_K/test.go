package main

import (
	"fmt"
)

func main() {
	date := 41
	result := findMinFibonacciNumbers(date)
	fmt.Printf("result is : %d", result)
}

func findMinFibonacciNumbers(k int) int {
	cemetery := []int{1, 1}
	result := 0
	for cemetery[len(cemetery)-1] > k {
		cemetery = append(cemetery, cemetery[len(cemetery)-1]+cemetery[len(cemetery)-2])
	}
	for i := len(cemetery) - 1; k > 0; i-- {
		if k >= cemetery[i] {
			k -= cemetery[i]
			result++
		}
	}
	return result
}
