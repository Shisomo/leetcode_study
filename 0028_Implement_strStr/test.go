package main

import "fmt"

func main() {
	data := "hello"
	dataN := "ll"
	result := strStr(data, dataN)
	fmt.Println(result)
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := range haystack {
		if haystack[i] != needle[0] {
			continue
		} else if len(haystack)-i >= len(needle) {
			for k := 0; k < len(needle); k++ {
				if haystack[i+k] != needle[k] {
					break
				}
				if k == len(needle)-1 {
					return i
				}
			}
		}
	}
	return -1
}
