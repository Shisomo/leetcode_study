package main

import "fmt"

func main() {
	result := numWaterBottles(13, 2)
	fmt.Printf("result is : %d", result)
}
func numWaterBottles(numBottles int, numExchange int) int {
	numdrink := numBottles
	for numBottles > 0 {
		chushu := numBottles / numExchange
		yushu := numBottles % numExchange
		numdrink = numdrink + chushu
		numBottles = chushu + yushu
		if numBottles < numExchange {
			// 增加可借瓶子选项
			//if numBottles+1 == numExchange {
			//	numBottles = numBottles + 1
			//} else {
			//	numBottles = 0
			//}
			numBottles = 0
		}
	}
	return numdrink
}
