package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:  7,
			Left: nil,
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(convertBST(root))
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Right)
	sum = sum + root.Val
	root.Val = sum
	traverse(root.Left)
}
