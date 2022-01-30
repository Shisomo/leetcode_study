package main

import (
	"fmt"
)

func main() {
	data := "hello world"
	result := reverseWords(data)
	fmt.Println(result)
}

func reverseWords(s string) string {
	a, b := 0, 0
	var cemetery []byte
	for a < len(s) {
		if s[a] == ' ' {
			for i := 1; i <= b; i++ {
				cemetery = append(cemetery, s[a-i])
			}
			b = 0
			cemetery = append(cemetery, ' ')
		} else if a+1 == len(s) {
			for i := 0; i <= b; i++ {
				cemetery = append(cemetery, s[a-i])
			}
		} else {
			b++
		}
		a++
	}
	return string(cemetery)
}
