package main

/*
	给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
	高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
	https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
*/

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	tt := CreateTree(arr)
	printTree(tt)
	fmt.Println()
	fmt.Println((*tt).Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	return &TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBST(nums[:middle]),
		Right: sortedArrayToBST(nums[middle+1:]),
	}
}

var i int = -1

func CreateTree(arr []int) *TreeNode {
	i++
	if i >= len(arr) {
		return nil
	}
	t := new(TreeNode)

	if arr[i] != 0 {
		t.Val = arr[i]
		t.Left = CreateTree(arr)
		t.Right = CreateTree(arr)
	}
	return t
}
func printTree(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Printf(" %d -> ", tree.Val)
	printTree(tree.Left)
	printTree(tree.Right)
}
