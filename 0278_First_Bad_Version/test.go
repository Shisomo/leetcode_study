package main

import (
	"fmt"
)

func main() {
	data := 5
	result := firstBadVersion(data)
	fmt.Printf("result is : %d", result)
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left, right := 1, n
	result := 0
	for left <= right {
		result = (left + right) / 2
		if !isBadVersion(result-1) && isBadVersion(result) {
			return result
		} else if isBadVersion(result-1) && isBadVersion(result) {
			right = result - 1
		} else {
			left = result + 1
		}
	}
	return -1
}
func isBadVersion(num int) bool {
	return num < 4
}
