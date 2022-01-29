package main

import (
	"fmt"
)

func main() {
	testdata1 := []int{1, 2, 3}
	testdata2 := []int{2, 5, 6}
	result := merge(testdata1, 3, testdata2, 3)
	fmt.Printf("result is : %d", result)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var result []int
	p, q := 0, 0
	for {
		if p == m {
			result = append(result, nums2[q:]...)
			break
		}
		if q == n {
			result = append(result, nums1[p:]...)
			break
		}
		if nums1[p] > nums2[q] {
			result = append(result, nums2[q])
			q++
		} else {
			result = append(result, nums1[p])
			p++
		}
	}
	copy(nums1, result)
	return result
}
