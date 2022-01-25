package main

import "fmt"

func main() {
	testdata := []int{2, 2, 2, 2, 2}
	result := plusOne(testdata)
	fmt.Printf("result is : %d", result)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}
