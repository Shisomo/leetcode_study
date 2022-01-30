package main

import (
	"fmt"
)

func main() {
	data := []int{2, 3, 5, 13, 15, 24, 33, 44, 53}
	result := singleNumber(data)
	fmt.Println(result)
}

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
