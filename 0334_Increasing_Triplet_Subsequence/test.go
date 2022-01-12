package main

import (
	"fmt"
	"math"
)

func main() {
	data := []int{1, 2, 3}
	result := increasingTriplet(data)
	fmt.Printf("result is : %v", result)
}
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] > second {
			return true
		} else if nums[i] <= first {
			first = nums[i]
		} else {
			second = nums[i]
		}
	}
	return false
}
