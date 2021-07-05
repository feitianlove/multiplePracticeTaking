package main

/*
	给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

	输入：n = 3
	输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

	链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii

*/
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return buildTree(1, n)
}

func buildTree(start, end int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if start >= end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		leftTreeNode := buildTree(start, i-1)
		rightTreeNode := buildTree(i+1, end)
		for _, left := range leftTreeNode {
			for _, right := range rightTreeNode {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return res
}
