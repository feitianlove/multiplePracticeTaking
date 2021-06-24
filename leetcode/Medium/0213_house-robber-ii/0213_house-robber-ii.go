package main

import "fmt"

/*
	你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
	同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。给定一个代表每个房屋存放金额的非负整数数组，
	计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

	输入：nums = [2,3,2]
	输出：3
	解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

	链接：https://leetcode-cn.com/problems/house-robber-ii
*/

func main() {
	nums := []int{1, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	o1 := dp(nums, 0, len(nums)-1)
	o2 := dp(nums, 1, len(nums))
	return max(o1, o2)
}
func dp(nums []int, start, end int) int {
	dp_i_1, dp_i_2 := 0, 0
	dp_i := 0
	for i := end - 1; i >= start; i-- {
		dp_i = max(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
