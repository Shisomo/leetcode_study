package main

import (
	"fmt"
)

func main() {
	data := ListNode{}
	result := reverseList(&data)
	fmt.Println(result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	new := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return new
}
