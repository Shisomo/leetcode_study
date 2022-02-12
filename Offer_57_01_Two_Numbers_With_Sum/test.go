package main

import (
	"fmt"
)

func main() {
	data := []int{14, 15, 16, 22, 53, 60}
	result := twoSum(data, 76)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return nil
}
