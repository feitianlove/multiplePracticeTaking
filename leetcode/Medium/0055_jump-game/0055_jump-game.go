package main

import "fmt"

/*

	给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	判断你是否能够到达最后一个下标。

	输入：nums = [2,3,1,1,4]
	输出：true
	解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

	https://leetcode-cn.com/problems/jump-game/
*/

func main() {
	fmt.Println(canJump1([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump1([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump1([]int{2, 0, 0}))
	fmt.Println(canJump1([]int{0}))

}

//贪心
func canJump2(nums []int) bool {
	j := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= j {
			j = i
		}
	}
	return j <= 0
}

// 动态规划
func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		flag := false
		for j := 0; j < i; j++ {
			if dp[j] && nums[j]+j >= i {
				flag = true
				break
			}
		}
		dp[i] = flag
	}
	return dp[len(nums)-1]
}

// 贪心
func canJump1(nums []int) bool {
	var farthest int = 0
	for i := 0; i < len(nums); i++ {
		farthest = max(farthest, nums[i]+i)
		if farthest <= i && i != len(nums)-1 {
			return false
		}
	}
	return farthest >= len(nums)-1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
