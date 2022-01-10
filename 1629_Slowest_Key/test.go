package main

import (
	"fmt"
)

func main() {
	testData := []int{2, 4, 6, 78, 98}
	testData2 := "abcdefg"
	result := slowestKey(testData, testData2)
	fmt.Printf("result is : %d", result)
}

// 贪心 前一天与后一天有点麻烦
func slowestKey(releaseTimes []int, keysPressed string) byte {
	cemetery := releaseTimes[0]
	result := keysPressed[0]
	for i, r := range releaseTimes {
		if i != len(releaseTimes)-1 {
			if releaseTimes[i+1]-r > cemetery || (releaseTimes[i+1]-r == cemetery && keysPressed[i+1] > result) {
				cemetery = releaseTimes[i+1] - r
				result = keysPressed[i+1]
			}
		}
	}
	return result
}
