package main

import "fmt"

func main() {
	testdata := []int{2, 2, 2, 2, 2}
	result := canJump(testdata)
	fmt.Printf("result is : %d", result)
}
func canJump(nums []int) bool {
	// sb程序改死人了
	cemetery := nums[0]
	if len(nums) == 1 {
		return true
	}
	if cemetery == 0 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		if cemetery >= i {
			if cemetery < i+nums[i] {
				cemetery = i + nums[i]
			}
			if cemetery >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
