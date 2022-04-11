package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 13}
	result := exchange(data)
	fmt.Println(result)
}

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 == 0 && nums[j]%2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		if nums[i]%2 == 1 {
			i++
		}
		if nums[j]%2 == 0 {
			j--
		}
	}
	return nums
}
