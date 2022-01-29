package main

import (
	"fmt"
)

func main() {
	data := []byte{3, 2, 2}
	reverseString(data)
	fmt.Printf("result is : %v", data)
}

func reverseString(s []byte) {
	var cemetery byte
	n := len(s)
	for i := 0; i < n/2; i++ {
		cemetery = s[i]
		s[i] = s[n-i-1]
		s[n-i-1] = cemetery
	}
}
