package main

import (
	"fmt"
	"math"
)

/*
	给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
	设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

	你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

	输入: [1,2,3,0,2]
	输出: 3
	解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
	链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown

*/
func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit1(prices))

}

func maxProfit(prices []int) int {
	var n = len(prices)
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		if i == 1 {
			dp[i][1] = max(dp[i-1][1], 0-prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		}
	}
	return dp[len(prices)-1][0]
}

//优化
func maxProfit1(prices []int) int {
	var n = len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt32
	pre_dp := 0
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, pre_dp-prices[i])
		pre_dp = temp
	}
	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
