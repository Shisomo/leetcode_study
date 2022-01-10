package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 13, 15, 24, 33, 44, 53}
	result := search(data, 3)
	fmt.Println(result)
}

func search(nums []int, target int) int {
	result := 0
	for _, n := range nums {
		if n == target {
			result++
		} else if n > target {
			break
		}
	}
	return result
}
