package main

import "fmt"

/*
	给定一个二叉树，检查它是否是镜像对称的。
	例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
	但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
	https://leetcode-cn.com/problems/symmetric-tree/
*/
func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 3

	root.Left = &left
	root.Right = &right

	fmt.Println(isSymmetric(&root))
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
	return recur(root.Left, root.Right)
}

func recur(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return recur(left.Right, right.Left) && recur(left.Left, right.Right)
}

// 不用递归
func isSymmetric2(root *TreeNode) bool {
	leftQ := make([]*TreeNode, 0)
	rightQ := make([]*TreeNode, 0)
	leftQ = append(leftQ, root)
	rightQ = append(rightQ, root)
	for len(leftQ) != 0 && len(rightQ) != 0 {
		leftCur := leftQ[0]
		rightCur := rightQ[0]
		leftQ = leftQ[1:]
		rightQ = rightQ[1:]
		if leftCur == nil && rightCur == nil {
			continue
		} else if leftCur != nil && rightCur != nil && leftCur.Val == rightCur.Val {
			leftQ = append(leftQ, leftCur.Left, leftCur.Right)
			rightQ = append(rightQ, rightCur.Right, rightCur.Left)
		} else {
			return false
		}

	}
	if len(leftQ) == 0 && len(rightQ) == 0 {
		return true
	} else {
		return false
	}
}
