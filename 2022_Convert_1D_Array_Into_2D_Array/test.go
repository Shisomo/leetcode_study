package main

import (
	"fmt"
)

func main() {
	testdata := []int{1, 2, 3, 4}
	result := construct2DArray(testdata, 2, 2)
	fmt.Printf("result is : %d", result)
}
func construct2DArray(original []int, m int, n int) [][]int {
	cemetery := make([][]int, 0, m)
	if m*n != len(original) {
		return cemetery
	}
	for i := 0; i < len(original); i += n {
		cemetery = append(cemetery, original[i:i+n])
	}
	return cemetery
}
