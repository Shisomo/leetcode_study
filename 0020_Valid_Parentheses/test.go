package main

import "fmt"

func main() {
	testString := "()(){([])}"
	result := isValid(testString)
	fmt.Printf("result is : %v", result)
}
func isValid(s string) bool {
	// 堆栈
	cemetery := []byte{}
	brackets := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		if len(cemetery)<1||(cemetery[len(cemetery)-1]!=brackets[s[i]]){
			cemetery = append(cemetery,s[i])
		}else {
			cemetery = cemetery[:len(cemetery)-1]
		}
	}
	if len(cemetery)==0{
		return true
	}else {
		return false
	}
}
