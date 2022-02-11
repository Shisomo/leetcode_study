package main

import (
	"fmt"
	"sort"
)

func main() {
	testdata := []int{1, 2, 3, 4}
	result := minimumDifference(testdata, 3)
	fmt.Println(result)
}
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	result := 100000
	for i := 0; i <= len(nums)-k; i++ {
		result = min(result, nums[i+k-1]-nums[i])
	}
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
