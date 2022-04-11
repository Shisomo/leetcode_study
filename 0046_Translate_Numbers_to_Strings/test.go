package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := 123
	res := translateNum(data)
	fmt.Println(res)
}

func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1


	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}
