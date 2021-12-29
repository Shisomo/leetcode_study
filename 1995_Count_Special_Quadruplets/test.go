package main

import (
	"fmt"
)

func main() {
	testdata := []int{1, 2, 3, 4}
	result := countQuadruplets(testdata)
	fmt.Printf("result is : %d", result)
}
func countQuadruplets(nums []int) int {
	cemetery := 0
	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			for c := b + 1; c < len(nums); c++ {
				for d := c + 1; d < len(nums); d++ {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						cemetery++
					}
				}
			}
		}
	}
	return cemetery
}
