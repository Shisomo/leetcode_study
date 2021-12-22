package main

import (
	"fmt"
)

func main() {
	houses := []int{1, 2, 3}
	result := rob(houses)
	fmt.Printf("result is : %d", result)
}
func rob(nums []int) int {
	p, q, r := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		p = q
		q = r
		if nums[i]+p > q {
			r = nums[i] + p
		} else {
			r = q
		}
	}
	return r
}
