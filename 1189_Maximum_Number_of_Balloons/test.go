package main

import "fmt"

func main() {
	testString := "ballonoballlioielfabnjmdnfasnefalkjfklejaflwk"
	result := maxNumberOfBalloons(testString)
	fmt.Println(result)
}
func maxNumberOfBalloons(text string) int {
	cemetery := [5]int{}
	for _, t := range text {
		if t == 'b' {
			cemetery[0]++
		}
		if t == 'a' {
			cemetery[1]++
		}
		if t == 'l' {
			cemetery[2]++
		}
		if t == 'o' {
			cemetery[3]++
		}
		if t == 'n' {
			cemetery[4]++
		}
	}
	cemetery[2] /= 2
	cemetery[3] /= 2
	res := 10000
	for _, ce := range cemetery {
		if ce < res {
			res = ce
		}
	}
	return res
}
