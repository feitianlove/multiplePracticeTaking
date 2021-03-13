package main

import "fmt"

/*
	给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
	https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
*/
func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	test1 := TreeNode{Val: 4}
	test2 := TreeNode{Val: 5}
	test3 := TreeNode{Val: 6}
	test4 := TreeNode{Val: 7}
	root.Left = &rootfirst
	root.Right = &rootSecond
	rootfirst.Left = &test1
	rootfirst.Right = &test2

	rootSecond.Left = &test3
	rootSecond.Right = &test4

	fmt.Println(levelOrderBottom(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	level := 0
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	orderBottom(root, level, &result)
	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}

func orderBottom(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}
	if len(*result) > level {
		(*result)[level] = append((*result)[level], root.Val)
	} else {
		*result = append(*result, []int{root.Val})
	}
	orderBottom(root.Left, level+1, result)
	orderBottom(root.Right, level+1, result)
}
func printTree(tree *TreeNode) {
	if tree == nil {
		return
	}
	printTree(tree.Left)
	printTree(tree.Right)
	fmt.Printf(" %d -> ", tree.Val)
}
