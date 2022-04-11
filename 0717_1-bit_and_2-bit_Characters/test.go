package main

import (
	"fmt"
)

func main() {
	n := []int{1, 1, 1, 0}
	result := isOneBitCharacter(n)
	fmt.Println(result)
}

func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2
		} else if bits[i] == 0 {
			i++
		}
	}
	return i == len(bits)-1
}
