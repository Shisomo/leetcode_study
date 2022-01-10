package main

import (
	"fmt"
)

func main() {
	data := []int{2, 3, 4, 5, 6, 1, 2}
	result := minArray(data)
	fmt.Println(result)
}

func minArray(numbers []int) int {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}
