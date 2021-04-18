package main

import (
	"fmt"
	"math"
)

/*
	给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
	你可以认为每种硬币的数量是无限的。

	输入：coins = [1, 2, 5], amount = 11
	输出：3
	解释：11 = 5 + 5 + 1

	链接：https://leetcode-cn.com/problems/coin-change

*/

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

// 动态规划
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for j := 0; j < len(coins); j++ {
			if i < coins[j] || dp[i-coins[j]] == -1 {
				continue
			}
			if dp[i] == -1 || dp[i] > dp[i-coins[j]]+1 {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}
	return dp[amount]
}

// 递归的over
func coinChange2(coins []int, amount int) int {
	var res = math.MaxInt32
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	for _, coin := range coins {
		if amount-coin < 0 {
			continue
		}
		subProblem := coinChange(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = min(res, 1+subProblem)
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
