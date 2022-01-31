package main

import (
	"fmt"
)

func main() {
	data := 25
	result := numberOfSteps(data)
	fmt.Println(result)
}

func numberOfStepsStudy(num int) int {
	result := 0
	for ; num > 0; num >>= 1 {
		result += num & 1
		if num > 1 {
			result++
		}
	}
	return result
}
func numberOfSteps(num int) int {
	result := 0
	for num != 0 {
		if num%2 == 0 {
			result++
		} else if num != 1 {
			result += 2
		} else {
			result++
		}
		num = num / 2
	}
	return result
}
