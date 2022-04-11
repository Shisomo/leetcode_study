package main

import "fmt"

func main() {
	var data TreeNode
	result := levelOrder(&data)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	var node *TreeNode
	for len(queue) != 0 {
		size := len(queue)
		var stack []int
		for i := 0; i < size; i++ {
			node = queue[0]
			stack = append(stack, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, stack)
	}
	return res
}
