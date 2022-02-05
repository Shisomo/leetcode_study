package main

import "fmt"

func main() {
	data := [][]int{{2, 3, 4, 5}}
	result := findNumberIn2DArray(data, 6)
	fmt.Println(result)
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}
