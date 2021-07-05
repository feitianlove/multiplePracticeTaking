package main

import (
	"math"
)

/*
	给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：
	二叉树的根是数组 nums 中的最大元素。
	左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
	右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
	返回有给定数组 nums 构建的 最大二叉树 。

	输入：nums = [3,2,1,6,0,5]
	输出：[6,3,5,null,2,0,null,null,1]
	解释：递归调用如下所示：
	- [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
		- [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
			- 空数组，无子节点。
			- [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
				- 空数组，无子节点。
				- 只有一个元素，所以子节点是一个值为 1 的节点。
		- [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
			- 只有一个元素，所以子节点是一个值为 0 的节点。
			- 空数组，无子节点。



	链接：https://leetcode-cn.com/problems/maximum-binary-tree

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums))
}

// 1 3 2
func build(nums []int, low, height int) *TreeNode {
	if len(nums[low:height]) == 0 {
		return nil
	}
	index := math.MinInt32
	max := math.MinInt32
	for i := low; i < height; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	var root *TreeNode = &TreeNode{
		Val: max,
	}
	// [)的区间不用减1
	//[0,1） [2,3)
	root.Left = build(nums, low, index)
	root.Right = build(nums, index+1, height)
	return root
}
