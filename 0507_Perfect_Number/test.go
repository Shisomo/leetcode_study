package main

import (
	"fmt"
)

func main() {
	n := 23
	result := checkPerfectNumber(n)
	fmt.Printf("result is : %v", result)
}

// func checkPerfectNumber(num int) bool {
//     // 太牛逼了
//     return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
// }
func checkPerfectNumber(num int) bool {
	cemetery := 0
	for i := 1; i*i < num; i++ {
		if num%i == 0 {
			cemetery = cemetery + i + num/i
		}
	}
	cemetery = cemetery - num
	return num == cemetery
}
