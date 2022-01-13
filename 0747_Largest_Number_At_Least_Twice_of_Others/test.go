package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{8, 10, 12}
	result := dominantIndex(n)
	fmt.Printf("result is : %v", result)
}

func dominantIndex(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	result, cemetery := -1, 0
	for i, n := range nums {
		if n > cemetery {
			cemetery = n
			result = i
		}
	}
	sort.Ints(nums)
	if nums[len(nums)-1]/2 >= nums[len(nums)-2] {
		return result
	} else {
		return -1
	}
}
