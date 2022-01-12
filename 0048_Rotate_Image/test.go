package main

import (
	"fmt"
)

func main() {
	data := [][]int{}
	rotate(data)
	fmt.Println(data)
}

func rotate(matrix [][]int) {
	n := len(matrix[0]) - 1
	ex := 0
	var exList []int
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			ex = matrix[i][j]
			matrix[i][j] = matrix[n-j][n-i]
			matrix[n-j][n-i] = ex
		}
	}
	for i := 0; i <= n/2; i++ {
		exList = matrix[i]
		matrix[i] = matrix[n-i]
		matrix[n-i] = exList
	}
}
