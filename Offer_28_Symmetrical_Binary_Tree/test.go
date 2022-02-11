package main

import "fmt"

func main() {
	var data TreeNode
	result := isSymmetric(&data)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return sym(root.Left, root.Right)
}
func sym(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return sym(left.Left, right.Right) && sym(left.Right, right.Left)
}
