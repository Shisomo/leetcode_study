package main

import (
	"fmt"
	"sort"
)

func main() {
	houses := []int{1, 2, 3}
	heaters := []int{1}
	result := findRadius(houses, heaters)
	fmt.Printf("result is : %v", result)
}
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	r := 0
	j := 0
	for _, house := range houses {
		dis := abs(house - heaters[j])
		for j+1 < len(heaters) && abs(house-heaters[j]) >= abs(house-heaters[j+1]) {
			j++
			if abs(house-heaters[j]) < dis {
				dis = abs(house - heaters[j])
			}
		}
		if dis > r {
			r = dis
		}
	}
	return r
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
