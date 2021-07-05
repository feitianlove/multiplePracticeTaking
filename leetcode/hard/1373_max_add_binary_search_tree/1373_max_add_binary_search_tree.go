package _1373_max_add_binary_search_tree

import "math"

/*

	给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。

	二叉搜索树的定义如下：

	任意节点的左子树中的键值都 小于 此节点的键值。
	任意节点的右子树中的键值都 大于 此节点的键值。
	任意节点的左子树和右子树都是二叉搜索树。


	输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
	输出：20
	解释：键值为 3 的子树是和最大的二叉搜索树。

*/
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {
	traverse(root)
	return maxSum
}

var maxSum int

func traverse(root *TreeNode) []int {
	if root == nil {
		return []int{1, math.MaxInt32, math.MinInt32, 0}
	}
	left := traverse(root.Left)
	right := traverse(root.Right)
	// 说明是left和right 是 bst树
	//res[0] 记录以 root 为根的二叉树是否是 BST，若为 1 则说明是 BST，若为 0 则说明不是 BST；
	//res[1] 记录以 root 为根的二叉树所有节点中的最小值；
	//res[2] 记录以 root 为根的二叉树所有节点中的最大值；
	//res[3] 记录以 root 为根的二叉树所有节点值之和。

	res := make([]int, 4)
	if left[0] == 1 && right[0] == 1 {
		if root.Val > left[2] || root.Val < right[1] {
			res[0] = 1
			res[1] = min(left[1], root.Val)
			res[2] = max(right[2], root.Val)
			res[3] = root.Val + left[3] + right[3]
			maxSum = max(maxSum, res[3])
		} else {
			res[0] = 0
		}
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
