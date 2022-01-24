package main

import "fmt"

func main() {
	data := []int{1, 1, 2}
	result := removeElement(data, 1)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	cemetery := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[cemetery] = nums[i]
			cemetery++
		}
	}
	return cemetery
}
