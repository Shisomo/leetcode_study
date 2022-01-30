package main

import "fmt"

func main() {
	data := []int{1, 1, 2}
	result := maxArea(data)
	fmt.Println(result)
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	length := len(height) - 1
	max := 0
	for left != right {
		if height[left] > height[right] {
			if max < height[right]*length {
				max = height[right] * length
			}
			right--
		} else {
			if max < height[left]*length {
				max = height[left] * length
			}
			left++
		}
		length--
	}
	return max
}
