package main

import "fmt"

/*
	给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

	一般来说，删除节点可分为两个步骤：

	首先找到需要删除的节点；
	如果找到了，删除它。
	说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

	示例:

	root = [5,3,6,2,4,null,7]
	key = 3

		5
	   / \
	  3   6
	 / \   \
	2   4   7

	给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。


	链接：https://leetcode-cn.com/problems/delete-node-in-a-bst

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}
	traverse(root)
	dd := deleteNode(root, 3)
	fmt.Println("======")
	//traverse(dd)
	fmt.Println(dd, dd.Left, dd.Right)
	fmt.Println(dd.Left.Left, dd.Left.Right)

}
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	if key == root.Val {
		//找到了分三种情况,没有子节点
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left != nil && root.Right == nil {
			root = root.Left
		} else if root.Right != nil && root.Left == nil {
			root = root.Right
		} else {
			//两个子节点都不是nil,找right最小的,
			temp := getMinNode(root.Right)
			root.Val = temp.Val
			root.Right = deleteNode(root.Right, root.Val)

		}
	}
	return root
}

//
func getMinNode(head *TreeNode) *TreeNode {
	for head.Left != nil {
		head = head.Left
	}
	return head
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	traverse(root.Left)
	traverse(root.Right)
}
