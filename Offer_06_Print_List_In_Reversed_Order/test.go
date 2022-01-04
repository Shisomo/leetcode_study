package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reversePrint(head *ListNode) []int {
//	result := []int{}
//	for head != nil {
//		result = append(result, head.Val)
//		head = head.Next
//	}
//	ex := 0
//	a := len(result)
//	for i := 0; i < a/2; i++ {
//		ex = result[i]
//		result[i] = result[a-i-1]
//		result[a-i-1] = ex
//	}
//	return result
//}
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)
}
func main() {
	var elementTest *ListNode
	reversePrint(elementTest)
}
