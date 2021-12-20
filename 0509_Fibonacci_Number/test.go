package main

import (
	"fmt"
)

func main() {
	n := 2
	result := fib(n)
	fmt.Printf("result is : %d", result)
}

func fib(n int) int {
	if n > 1 {
		cemetery := 0
		left := 0
		right := 1
		for i := 1; i < n; i++ {
			cemetery = left + right
			left = right
			right = cemetery
		}
		return cemetery
	} else {
		return n
	}
}
