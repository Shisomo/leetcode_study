package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 3}
	result := containsDuplicate(data)
	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {
	cemetery := make(map[int]int)
	for _, n := range nums {
		cemetery[n]++
	}
	for _, v := range cemetery {
		if v > 1 {
			return true
		}
	}
	return false
}
