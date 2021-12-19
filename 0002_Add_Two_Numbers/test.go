package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			}},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			}},
	}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Printf("%d", result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	head := &ListNode{Val: 0, Next: nil}
	caculate := head
	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}
		caculate.Next = &ListNode{Val: (x + y + carry) % 10, Next: nil}
		caculate = caculate.Next
		carry = (x + y + carry) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if carry > 0 {
			caculate.Next = &ListNode{Val: carry, Next: nil}
		}
	}
	return head.Next
}
