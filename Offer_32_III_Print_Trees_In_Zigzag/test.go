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
	level := 1
	for len(queue) != 0 {
		size := len(queue)
		var curLevel []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level%2 == 1 {
			res = append(res, curLevel)
		} else {
			res = append(res, reverseArray(curLevel))
		}
		level++
	}
	return res
}

func reverseArray(array []int) []int {
	n := len(array)
	for i := 0; i < n/2; i++ {
		temp := array[i]
		array[i] = array[n-1-i]
		array[n-1-i] = temp
	}
	return array
}
