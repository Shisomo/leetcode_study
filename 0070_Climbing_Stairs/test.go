package main

import "fmt"

func main() {
	testString := 4
	result := climbStairs(testString)
	fmt.Printf("result is : %d", result)
}
func climbStairs(n int) int {
	p := 0
	q := 0
	result := 1
	for Step := 1; Step <= n; Step++ {
		p = q
		q = result
		result = p + q
	}
	return result
}
