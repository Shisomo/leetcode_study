package main

import (
	"fmt"
	"strconv"
)

func main() {
	date := 6
	result := simplifiedFractions(date)
	fmt.Println(result)
}

func simplifiedFractions(n int) []string {
	var result []string
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(j, i) == 1 {
				result = append(result, strconv.Itoa(j)+"/"+strconv.Itoa(i))
			}
		}
	}
	return result
}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
