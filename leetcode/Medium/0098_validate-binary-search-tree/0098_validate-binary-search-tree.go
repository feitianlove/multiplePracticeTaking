package main

import "fmt"

/*
	给定一个二叉树，判断其是否是一个有效的二叉搜索树。
	假设一个二叉搜索树具有如下特征：

	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。

	链接：https://leetcode-cn.com/problems/validate-binary-search-tree
*/
func main() {
	root := &TreeNode{
		Val:  10,
		Left: &TreeNode{Val: 5},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 20},
		},
	}
	fmt.Println(isValidBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return traverse(root, nil, nil)
}

func traverse(root *TreeNode, minTreeNode *TreeNode, maxTreeNode *TreeNode) bool {
	if root == nil {
		return true
	}
	if minTreeNode != nil && root.Val <= minTreeNode.Val {
		return false
	}
	if maxTreeNode != nil && root.Val >= maxTreeNode.Val {
		return false
	}
	return traverse(root.Left, minTreeNode, root) && traverse(root.Right, root, maxTreeNode)

}
