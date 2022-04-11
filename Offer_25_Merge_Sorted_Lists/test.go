package main

import "fmt"

func main() {
	data := new(ListNode)
	result := mergeTwoLists(data, data)
	fmt.Println(result)
}

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			head.Next = l1
			head = head.Next
			l1 = l1.Next
		} else {
			head.Next = l2
			head = head.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return p.Next
}
