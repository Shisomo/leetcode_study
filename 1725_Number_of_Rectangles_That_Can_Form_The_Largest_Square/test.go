package main

import (
	"fmt"
)

func main() {
	testdata := [][]int{{3, 5}, {5, 14}, {5, 14}, {5, 14}}
	result := countGoodRectangles(testdata)
	fmt.Printf("result is : %d", result)
}

func countGoodRectangles(rectangles [][]int) int {
	result, max, minre := 0, 0, 0
	for _, rect := range rectangles {
		minre = min(rect[0], rect[1])
		if minre > max {
			result = 1
			max = minre
		} else if minre == max {
			result++
		}
	}
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
