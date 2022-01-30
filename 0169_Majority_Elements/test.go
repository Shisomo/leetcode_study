package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 13, 15, 24, 33, 44, 53}
	result := majorityElement(data)
	fmt.Println(result)
}

func majorityElement(nums []int) int {
	cemetery := make(map[int]int)
	for _, n := range nums {
		cemetery[n]++
	}
	for i, c := range cemetery {
		if c > len(nums)/2 {
			return i
		}
	}
	return 0
}
