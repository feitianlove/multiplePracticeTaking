package main

/*

	给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。



	输入：root = [3,1,4,null,2], k = 1
	输出：1

	https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
*/
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var rank int
var res int

func kthSmallest(root *TreeNode, k int) int {
	rank = 0
	res = 0
	traverse(root, k)
	return res
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	rank++
	if rank == k {
		res = root.Val
		return
	}
	traverse(root.Right, k)
}
