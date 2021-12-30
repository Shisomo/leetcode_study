package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{8, 10, 12}
	result := isNStraightHand(n, 3)
	fmt.Printf("result is : %v", result)
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	cemetery := make(map[int]int)
	for _, i := range hand {
		cemetery[i]++
	}
	for _, i := range hand {
		if cemetery[i] == 0 {
			continue
		}
		for j := i; j < i+groupSize; j++ {
			print(j)
			if cemetery[j] == 0 {
				return false
			}
			cemetery[j]--
		}
	}
	return true
}
