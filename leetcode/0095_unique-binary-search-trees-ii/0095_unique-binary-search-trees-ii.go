package main

/*
	给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
	输入：3
	输出：
	[
	  [1,null,3,2],
	  [3,2,null,1],
	  [3,1,null,null,2],
	  [2,1,3],
	  [1,null,2,null,3]
	]
	解释：
	以上的输出对应以下 5 种不同结构的二叉搜索树：

	   1         3     3      2      1
		\       /     /      / \      \
		 3     2     1      1   3      2
		/     /       \                 \
	   2     1         2                 3

链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func generateTrees(n int) []*TreeNode {
	return nil
}
