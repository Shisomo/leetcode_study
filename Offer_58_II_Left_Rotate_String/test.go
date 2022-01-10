package main

import "fmt"

func main() {
	data := "lskdjfls"
	dataNum := 3
	result := reverseLeftWords(data, dataNum)
	fmt.Println(result)
}

func reverseLeftWords(s string, n int) string {
	s_byte := []byte(s)
	s1 := s_byte[:n]
	s2 := s_byte[n:]
	result := append(s2, s1...)
	return string(result)
}
