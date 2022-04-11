package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 13, 15, 24, 33, 44, 53}
	result := maxProfit(data)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	/*
		给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
		设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
		注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	*/
	if len(prices) == 0 {
		return 0
	}
	min, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}
