package main

import "fmt"

func main() {
	data := 6
	result := numWays(data)
	fmt.Println(result)
}

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	first, second, result := 1, 1, 1
	for i := 1; i < n; i++ {
		result = (first + second) % 1000000007
		first = second
		second = result
	}
	return result
}
