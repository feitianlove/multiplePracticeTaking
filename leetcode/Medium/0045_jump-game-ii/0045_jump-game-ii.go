package main

import (
	"fmt"
	"math"
)

/*
	给你一个非负整数数组 nums ，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。你的目标是使用最少的跳跃次数到达
	数组的最后一个位置。
	假设你总是可以到达数组的最后一个位置。


	输入: nums = [2,3,1,1,4]
	输出: 2
	解释: 跳到最后一个位置的最小跳跃数是 2。
	从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。


	链接：https://leetcode-cn.com/problems/jump-game-ii
*/
func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump2(nums))
}

// 动态规划
func jump(nums []int) int {
	var dp []int = make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

//贪心
//2, 3, 0, 1, 4
func jump2(nums []int) int {
	var farthest = 0
	var end = 0
	var res int
	for i := 0; i < len(nums); i++ {
		farthest = max(farthest, nums[i]+i)
		if end == i && i != len(nums)-1 {
			res++
			end = farthest
		}
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
