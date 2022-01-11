package main

import "fmt"

func main() {
	data := "abacdidkfjdkjfkd"
	result := firstUniqChar(data)
	fmt.Println(result)
}

func firstUniqChar(s string) byte {
	sByte := []byte(s)
	cemetery := make(map[byte]int)
	for _, c := range sByte {
		cemetery[c]++
	}
	for _, c := range sByte {
		if cemetery[c] == 1 {
			return c
		}
	}
	return ' '
}
