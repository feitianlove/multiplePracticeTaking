package main

import "fmt"

/*
	给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

	输入: [1, 5, 11, 5]
	输出: true
	解释: 数组可以分割成 [1, 5, 5] 和 [11].

	注意:
	每个数组中的元素不会超过 100
	数组的大小不会超过 200

	链接：https://leetcode-cn.com/problems/partition-equal-subset-sum



*/

func main() {
	fmt.Println(canPartition([]int{1, 2, 5, 2}))
	//fmt.Println(canPartition([]int{1, 2, 3, 5}))

}

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	// 经典背包问题
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= target; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}

		}
	}
	return dp[len(nums)][target]
}

//[]int{1, 2, 5, 2}
