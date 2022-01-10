package main

import "fmt"

func main() {
	data := []int{2, 3, 4, 5}
	result := findRepeatNumber(data)
	fmt.Println(result)
}

func findRepeatNumber(nums []int) int {
	cemetery := map[int]int{}
	for _, n := range nums {
		cemetery[n]++
		if cemetery[n] > 1 {
			return n
		}
	}
	return -1
}
