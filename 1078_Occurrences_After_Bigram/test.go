package main

import (
	"fmt"
	"strings"
)

func main() {
	data1, data2, data3 := "a", "dd", "df"
	result := findOcurrences(data1, data2, data3)
	fmt.Printf("result is : %s", result)
}

func findOcurrences(text string, first string, second string) []string {
	str2slice := strings.Split(text, " ")
	cemetery := []string{}
	for i := 0; i < len(str2slice)-2; i++ {
		if str2slice[i] == first && str2slice[i+1] == second {
			cemetery = append(cemetery, str2slice[i+2])
		}
	}
	return cemetery
}
