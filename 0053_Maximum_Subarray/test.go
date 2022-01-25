package main

import "fmt"

func main() {
	testdata := []int{2, 2, 2, 2, 2}
	result := maxSubArray(testdata)
	fmt.Printf("result is : %d", result)
}

func maxSubArray(nums []int) int {
	pst, result := 0, -10000
	for _, n := range nums {
		if pst < 0 {
			pst = n
		} else {
			pst += n
		}
		if pst > result {
			result = pst
		}
	}
	return result
}
