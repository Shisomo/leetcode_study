package main

import "fmt"

func main() {
	testdata := "hello world "
	result := lengthOfLastWord(testdata)
	fmt.Printf("result is : %d", result)
}

func lengthOfLastWord(s string) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			result++
		}
		if s[i] == ' ' && result > 0 {
			break
		}
	}
	return result
}
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 1
		} else {
			digits[i] += 1
			return digits
		}
	}
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}
