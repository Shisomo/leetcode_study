package main

import "fmt"

func main() {
	fmt.Printf(" ")
	a := []int{1, 2, 3, 4, 5, 6}
	b := 10
	c := sum(a, b)
	fmt.Printf("%d", c)
}
// 主要思路: 循环一遍, 判断当前数据的互补是否存在于墓场, 存在的话返回墓场值与当前值
func sum(nums []int, target int) []int {
	// cemetery 为墓场，存放数据进行比较
	cemetery := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		compare := target - nums[i]
		if _, ok := cemetery[compare]; ok {
			return []int{cemetery[compare], i}
		}
		cemetery[nums[i]] = i
	}
	return nil
}
