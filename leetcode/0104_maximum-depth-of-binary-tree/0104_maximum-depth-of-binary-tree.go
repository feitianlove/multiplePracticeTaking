package main

import (
	"fmt"
)

/*
	给定一个二叉树，找出其最大深度。
	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
	https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
*/
func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(maxDepth(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//if root.Left == nil && root.Right == nil {
	//	return left + 1
	//}
	//if root.Left != nil {
	//	left = 1 + left + maxDepth(root.Left)
	//}
	//if root.Right != nil {
	//	right = 1 + right + maxDepth(root.Right)
	//}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// 迭代
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			temp := queue[0]
			queue = queue[1:]
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		depth++
	}
	return depth
}
