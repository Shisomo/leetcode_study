package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{1, 2, 3, 4}
	result := numFriendRequests(n)
	fmt.Printf("result is : %d", result)
}

func numFriendRequests(ages []int) int {
	// 套两层循环的O(n^2)超时了
	// 此解法为学习官方解法
	// ages排序，寻找每个age满足的年龄区间
	sort.Ints(ages)
	cemetery := 0
	left, right := 0, 0
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		cemetery += right - left
	}
	return cemetery
}
