package main

/*
	给定一个二叉树，找出其最小深度。
	最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
	说明：叶子节点是指没有子节点的节点
	https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
*/

import "fmt"

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2
	right.Right = &TreeNode{Val: 1}
	root.Left = &left
	root.Right = &right
	res := minDepth(&root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil {
		return 1 + minDepth(root.Right)
	} else if root.Right == nil {
		return 1 + minDepth(root.Left)
	} else {
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}

}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// 回溯算法
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 1
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			temp := queue[0]
			queue = queue[1:]
			if temp.Left == nil && temp.Right == nil {
				return depth
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
		}
		depth++
	}
	return depth
}
