package main

import (
	"fmt"
)

type StockPrice struct {
	price map[int]int
}

func Constructor() StockPrice {
	var price map[int]int
	return StockPrice{price}
}

func (this *StockPrice) Update(timestamp int, price int) {

}

func (this *StockPrice) Current() int {
	this.price[0] = 1
}

func (this *StockPrice) Maximum() int {

}

func (this *StockPrice) Minimum() int {

}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */

func main() {
	timestamp := 4
	price := 3
	obj := Constructor()
	obj.Update(timestamp, price)
	param2 := obj.Current()
	param3 := obj.Maximum()
	param4 := obj.Minimum()
	fmt.Printf("result is : %v %v %v\n", param2, param3, param4)
}
