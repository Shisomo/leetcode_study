package main

import "fmt"

func main() {
	testString := "IV"
	result := romanToInt(testString)
	fmt.Printf("result is : %v", result)
}
func romanToInt(s string) int {
	exchangeRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && exchangeRoman[s[i]] < exchangeRoman[s[i+1]] {
			result = result - exchangeRoman[s[i]]
		} else {
			result = result + exchangeRoman[s[i]]
		}
	}
	return result
}
