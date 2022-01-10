package main

import "fmt"

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 7, 8, 9}
	result := missingNumber(data)
	fmt.Println(result)
}

// func missingNumber(nums []int) int {
//
// 	for i, n := range nums {
// 		if i != n {
// 			return i
// 		}
// 	}
// 	return len(nums)
// }

func missingNumber(nums []int) int {
	// 二分法
	i, j := 0, len(nums)-1
	n := 0
	for i <= j {
		n = (i + j) / 2
		if nums[n] == n {
			i = n + 1
		} else {
			j = n - 1
		}
	}
	return i
}
