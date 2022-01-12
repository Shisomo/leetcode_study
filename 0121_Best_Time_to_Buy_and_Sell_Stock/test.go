package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 13, 15, 24, 33, 44, 53}
	result := maxProfit(data)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
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
