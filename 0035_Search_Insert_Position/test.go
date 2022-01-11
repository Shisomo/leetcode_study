package main

import "fmt"

func main() {
	testdata := []int{2, 2, 2, 2, 2}
	result := searchInsert(testdata, 3)
	fmt.Printf("result is : %d", result)
}

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	left, right := 0, len(nums)-1
	result := 0
	for left <= right {
		result = (left + right) / 2
		if nums[result] == target {
			return result
		} else if nums[result] < target {
			left = result + 1
		} else {
			right = result - 1
		}
	}
	return (left+right)/2 + 1
}
