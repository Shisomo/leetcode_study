package main

import (
	"fmt"
)

func main() {
	testdata := 45
	result := totalMoney(testdata)
	fmt.Printf("result is : %d", result)
}

func totalMoney(n int) int {
	week := n / 7
	day := n % 7
	if week == 0 {
		return n * (n + 1) / 2
	}
	result := week*28 + (week-1)*week*7/2 + day*week + (1+day)*day/2
	return result
}
