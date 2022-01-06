package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "/ddfdfdf/sdfs/d/f/sd/f/s/f/s/"
	result := simplifyPath(testString)
	fmt.Printf("result is : %d", result)
}
func simplifyPath(path string) string {
	stack := []string{}
	for _, p := range strings.Split(path, "/") {
		if p == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if p != "" && p != "." {
			stack = append(stack, p)
		}
	}
	return "/" + strings.Join(stack, "/")
}
