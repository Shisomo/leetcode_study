package main

import "fmt"

func main() {
	data := []int{1, 1, 2}
	result := removeDuplicates(data)
	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	result := []int{}
	cemetery := map[int]int{}
	for _, n := range nums {
		if cemetery[n] == 0 {
			result = append(result, n)
		}
		cemetery[n]++
	}
	for i := range result {
		nums[i] = result[i]
	}
	return len(result)
}
