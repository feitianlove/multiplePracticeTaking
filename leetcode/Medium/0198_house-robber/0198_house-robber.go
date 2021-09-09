package main

import "fmt"

/*
	你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房
	屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

	输入：[2,7,9,3,1]
	输出：12
	解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。



	链接：https://leetcode-cn.com/problems/house-robber
*/
func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
	fmt.Println(rob3(nums))
}

// 加备忘录的递归
var mem []int

func rob(nums []int) int {
	mem = make([]int, len(nums))
	for i := 0; i < len(mem); i++ {
		mem[i] = -1
	}
	return dp(nums, 0)
}
func dp(nums []int, start int) int {
	for start >= len(nums) {
		return 0
	}
	if mem[start] != -1 {
		return mem[start]
	}
	//就两种情况抢和不抢，取最大值
	big := max(dp(nums, start+1), nums[start]+dp(nums, start+2))
	mem[start] = big
	return big
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//动规优化
func rob2(nums []int) int {
	dp := make([]int, len(nums)+2)
	for i := len(mem) - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}
	return dp[0]
}

//优化空间复杂度
func rob3(nums []int) int {
	dp_i_1, dp_i_2 := 0, 0
	dp_i := 0
	for i := len(nums) - 1; i >= 0; i-- {
		dp_i = max(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i

	}
	return dp_i
}
