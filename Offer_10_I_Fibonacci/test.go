package main

import "fmt"

func main() {
	data := 6
	result := fib(data)
	fmt.Println(result)
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a, b, c := 0, 0, 1
	for i := 1; i < n; i++ {
		a = b
		b = c
		c = (a + b) % 1000000007
	}
	return c
}
