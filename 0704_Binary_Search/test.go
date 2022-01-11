package main

import (
	"fmt"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 6
	result := search(data, n)
	fmt.Printf("result is : %v", result)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := 0
	for left <= right {
		result = (left + right) / 2
		if nums[result] < target {
			left = result + 1
		} else if nums[result] > target {
			right = result - 1
		} else {
			return result
		}
	}
	return -1
}
