package main

import "fmt"

func main() {
	data := new(ListNode)
	result := deleteNode(data, 5)
	fmt.Println(result)
}

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	root := head
	for root.Next != nil {
		if root.Next.Val == val {
			root.Next = root.Next.Next
			break
		}
		root = root.Next
	}
	return head
}
