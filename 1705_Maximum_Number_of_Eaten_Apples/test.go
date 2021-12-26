package main

import (
	"fmt"
)

func main() {
	testdata := []int{1, 2, 3, 4}
	testdata2 := []int{1, 2, 3, 4}

	result := eatenApples(testdata, testdata2)
	fmt.Printf("result is : %d", result)
}
// 贪心 前一天与后一天有点麻烦
func eatenApples(apples []int, days []int) int {
	cemetery := 0
	result := 0
	for i := 0; i < len(apples); i++ {
		return 0
	}
}

func max(compare_a int, compare_b int) int {
	if compare_a < compare_b {
		return compare_b
	} else {
		return compare_a
	}
}
