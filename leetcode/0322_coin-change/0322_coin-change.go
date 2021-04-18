package main

import (
	"fmt"
)

/*
	给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
	你可以认为每种硬币的数量是无限的。

	输入：coins = [1, 2, 5], amount = 11
	输出：3
	解释：11 = 5 + 5 + 1

	链接：https://leetcode-cn.com/problems/coin-change

	dp(n) = 0 , n = 0
	dp(n) = -1, n < 0
	dp(n) = min(dp(n), 1 + dp(n-coin))
*/

func main() {
	fmt.Println(coinChange2([]int{1, 2, 5}, 11))
	fmt.Println(coinChange2([]int{2}, 3))

	//fmt.Println(coinChange([]int{3}, 3))

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// dp
//var dp = make(map[int]int)
//
//func coinChange(coins []int, amount int) int {
//	var res = math.MaxInt32
//	if amount < 0 {
//		return -1
//	}
//	if amount == 0 {
//		return 0
//	}
//	if r, ok := dp[amount]; ok {
//		return r
//	}
//	for _, coin := range coins {
//
//		if amount-coin < 0 {
//			continue
//		}
//		subProblem := coinChange(coins, amount-coin)
//		if subProblem == -1 {
//			continue
//		}
//		res = min(res, 1+subProblem)
//	}
//	if res == math.MaxInt32 {
//		return -1
//	}
//	dp[amount] = res
//	return res
//}

func coinChange2(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i-coin < 0 || dp[i-coin] == -1 {
				continue
			}
			if dp[i] == -1 || dp[i] > dp[i-coin]+1 {
				dp[i] = 1 + dp[i-coin]
			}
		}
	}
	return dp[amount]
}
