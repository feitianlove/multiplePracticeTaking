package main

import "fmt"

/*
	翻转一棵二叉树。

	输入：

		 4
	   /   \
	  2     7
	 / \   / \
	1   3 6   9
	输出：

		 4
	   /   \
	  7     2
	 / \   / \
	9   6 3   1

	链接：https://leetcode-cn.com/problems/invert-binary-tree
*/

func main() {
	root := TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(root)
	fmt.Println(invertTree(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var temp *TreeNode
	temp = root.Left
	root.Left = root.Right
	root.Right = temp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
