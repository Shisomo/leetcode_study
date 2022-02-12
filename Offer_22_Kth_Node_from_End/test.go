package main

import "fmt"

func main() {
	data := new(ListNode)
	result := getKthFromEnd(data, 5)
	fmt.Println(result)
}

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
