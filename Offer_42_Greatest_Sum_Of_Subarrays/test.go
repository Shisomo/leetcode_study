package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 13, 15, 24, 33, 44, 53}
	result := maxSubArray(data)
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
