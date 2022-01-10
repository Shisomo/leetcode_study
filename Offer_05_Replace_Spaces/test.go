package main

import "fmt"

func main() {
	data := "Hello World !"
	result := replaceSpace(data)
	fmt.Println(result)
}

func replaceSpace(s string) string {
	Sb := []byte(s)
	result := make([]byte, 0)
	for _, sLoop := range Sb {
		if sLoop == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, sLoop)
		}
	}
	return string(result)
}
