package main

import (
	"fmt"
	"strconv"
)

/*
	给你一个整数数组 nums 和一个整数 target 。
	向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
	例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
	返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

	输入：nums = [1,1,1,1,1], target = 3
	输出：5
	解释：一共有 5 种方法让最终目标和为 3 。
	-1 + 1 + 1 + 1 + 1 = 3
	+1 - 1 + 1 + 1 + 1 = 3
	+1 + 1 - 1 + 1 + 1 = 3
	+1 + 1 + 1 - 1 + 1 = 3
	+1 + 1 + 1 + 1 - 1 = 3

	链接：https://leetcode-cn.com/problems/target-sum
*/
func main() {
	var nums = []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays2(nums, 3))
}

var dp map[string]int

func findTargetSumWays(nums []int, target int) int {
	dp = make(map[string]int)
	return dfs(nums, target, 0)
}

//回溯算法解
func dfs(nums []int, target int, index int) int {
	if index == len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}
	k := strconv.Itoa(index) + "," + strconv.Itoa(target)
	if value, ok := dp[k]; ok {
		return value
	}
	res := dfs(nums, target-nums[index], index+1) + dfs(nums, target+nums[index], index+1)
	dp[k] = res
	return res
}

//参考大佬的思维，把它变成背包类型的问题
//sum(A) - sum(B) = target
//sum(A) = target + sum(B)
//sum(A) + sum(A) = target + sum(B) + sum(A)
//2 * sum(A) = target + sum(nums)
//也就是说 num中存在多少个子集的和等于 （target + sum(nums)）/2
func findTargetSumWays2(nums []int, target int) int {
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum < target || (sum+target)%2 == 1 {
		return 0
	}
	return subsets(nums, (sum+target)/2)
}

// dp[i][j] 就是前n个物品装成j重量有多少中装法
func subsets(nums []int, target int) int {
	var dp = make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1

	}
	n := len(nums)
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			//装不下
			if nums[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				// 能装下
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][target]
}
