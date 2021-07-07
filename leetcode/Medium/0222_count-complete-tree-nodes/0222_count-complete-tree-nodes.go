package main

import "math"

/*

	给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
	完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
	若最底层为第 h 层，则该层包含 1~ 2h 个节点。



	示例 1：


	输入：root = [1,2,3,4,5,6]
	输出：6

	链接：https://leetcode-cn.com/problems/count-complete-tree-nodes
*/
// TODO 中文语境和英文语境似乎有点区别，我们说的完全二叉树对应英文 Complete Binary Tree，没有问题。
// TODO 但是我们说的满二叉树对应英文 Perfect Binary Tree，而英文中的 Full Binary Tree 是指一棵二叉树的所有节点要么没有孩子节点，
// TODO 要么有两个孩子节点。 我们说的满二叉树，是一种特殊的完全二叉树，每层都是是满的，像一个稳定的三角形

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// so easy 普通二叉树的通用解法
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

//满二叉树的解法
func countNodes2(root *TreeNode) int {
	h := 0
	for root != nil {
		root = root.Left
		h++
	}
	return int(math.Pow(2, float64(h)) - 1)
}

//综合解法
func countNodes3(root *TreeNode) int {
	l, r := root, root
	lh, rh := 0, 0
	for l != nil {
		l = l.Left
		lh++
	}
	for r != nil {
		r = r.Right
		rh++
	}
	if lh == rh {
		return int(math.Pow(2, float64(lh)) - 1)
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
