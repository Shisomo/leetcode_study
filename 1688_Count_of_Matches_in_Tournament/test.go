package main

import (
	"fmt"
)

func main() {
	testData := 3
	result := numberOfMatches(testData)
	fmt.Printf("result is : %d", result)
}

func numberOfMatches(n int) int {
	return n - 1
}
