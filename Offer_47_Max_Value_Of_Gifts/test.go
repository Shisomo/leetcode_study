package main

import "fmt"

func main() {
	data := [][]int{{2, 3, 5, 13}, {15, 24, 33, 44}}
	result := maxValue(data)
	fmt.Println(result)
}

func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				grid[i][j] = grid[i][j]
			} else if i == 0 && j != 0 {
				grid[i][j] += grid[i][j-1]
			} else if i != 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += max(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
