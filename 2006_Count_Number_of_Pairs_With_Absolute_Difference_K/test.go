package main

import (
	"fmt"
)

func main() {
	testdata1 := []int{2, 1, 1, 2}
	result := countKDifference(testdata1, 1)
	fmt.Printf("result is : %s", result)
}

func countKDifference(nums []int, k int) int {
	result := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]-nums[j] == k || nums[i]-nums[j] == -k {
				result++
			}
		}
	}
	return result
}
func countKDifferenceHush(nums []int, k int) int {
	result := 0
	cemetery := map[int]int{}
	for _, n := range nums {
		cemetery[n]++
		result += cemetery[n-k] + cemetery[n+k]

	}
	return result
}
