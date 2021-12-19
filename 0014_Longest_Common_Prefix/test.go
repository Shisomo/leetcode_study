package main

import "fmt"

func main() {
	testString := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(testString)
	fmt.Printf("result is : %v", result)
}
func longestCommonPrefix(strs []string) string {
	cemetery := strs[0]
	for i := 0; i < len(cemetery); i++ {
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != cemetery[i] {
				return cemetery[:i]
			}
		}
	}
	return cemetery
}
