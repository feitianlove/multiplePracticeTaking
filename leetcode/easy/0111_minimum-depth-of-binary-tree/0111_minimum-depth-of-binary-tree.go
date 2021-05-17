package main

import "fmt"

/*
	给定一个二叉树，找出其最小深度。
	最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
	说明：叶子节点是指没有子节点的节点。

	输入：root = [3,9,20,null,null,15,7]
	输出：2
	https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
*/
func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	right.Right = &TreeNode{Val: 3}
	left.Right = &TreeNode{Val: 3}
	root.Left = &left
	root.Right = &right
	res := minDepth(&root)
	fmt.Println(res, "res")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)

	deep := 1
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[0].Left == nil && queue[0].Right == nil {
				return deep
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		deep++
	}
	return deep
}
