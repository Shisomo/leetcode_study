package main

import "fmt"

func main() {
	testString := [][]int{{1, 2}, {1, 3}, {2, 3}}
	result := findJudge(3, testString)
	fmt.Printf("result is : %v", result)
}
func findJudge(n int, trust [][]int) int {
	if len(trust) < n-1 || len(trust) > (n-1)*(n-1) {
		return -1
	}
	in := make([]int, n+1)
	out := make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		in[trust[i][1]]++
		out[trust[i][0]]++
	}
	for i := 1; i <= n; i++ {
		if in[i] == n-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}
