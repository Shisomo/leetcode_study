package main

import (
	"fmt"
)

func main() {
	n := 2
	result := tribonacci(n)
	fmt.Printf("result is : %d", result)
}

func tribonacci(n int) int {
	if n > 2 {
		cemetery := 0
		one := 0
		two := 0
		three := 1
		for i := 2; i < n; i++ {
			cemetery = one + two + three
			one = two
			two = three
			three = cemetery
		}
		return cemetery
	}else if n == 2 {
		return 1
	}else {
		return n
	}
}
