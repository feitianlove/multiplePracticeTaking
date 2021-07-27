package main

import "fmt"

/*
	给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
	请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
	假设每一种面额的硬币有无限个。
	题目数据保证结果符合 32 位带符号整数。



	示例 1：

	输入：amount = 5, coins = [1, 2, 5]
	输出：4
	解释：有四种方式可以凑成总金额：
	5=5
	5=2+2+1
	5=2+1+1+1
	5=1+1+1+1+1

	链接：https://leetcode-cn.com/problems/coin-change-2
*/
func main() {
	var coins = []int{1, 2, 5}
	fmt.Println(change(5, coins))
}

func change(amount int, coins []int) int {
	var dp [][]int = make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}
