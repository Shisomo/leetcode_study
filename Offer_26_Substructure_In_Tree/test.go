package main

import "fmt"

func main() {
	var data, dataTest TreeNode
	result := isSubStructure(&data, &dataTest)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isSame(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)

}

func isSame(a, b *TreeNode) bool {
	if a == nil && b != nil {
		return false
	}
	if b == nil {
		return true
	}
	if a.Val == b.Val {
		return isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
	}
	return false
}
