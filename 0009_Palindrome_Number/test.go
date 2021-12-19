package main

import "fmt"

func main() {
	testString := 121
	result := isPalindrome(testString)
	fmt.Printf("result is : %v", result)
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x != 0 && x%10 == 0 {
		return false
	}
	leftNums := 0
	for leftNums < x {
		leftNums = leftNums*10 + x%10
		x = x / 10
	}
	if leftNums == x || leftNums/10 == x {
		return true
	} else {
		return false
	}
}
