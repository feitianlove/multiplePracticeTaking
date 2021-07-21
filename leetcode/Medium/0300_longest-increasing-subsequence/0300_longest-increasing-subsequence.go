package main

import "fmt"

/*
	给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
	子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

	输入：nums = [10,9,2,5,3,7,101,18]
	输出：4
	解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

	链接：https://leetcode-cn.com/problems/longest-increasing-subsequence

*/
func main() {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = 1
	var res int
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
