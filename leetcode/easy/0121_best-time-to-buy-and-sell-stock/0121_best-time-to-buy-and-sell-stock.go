package main

import (
	"fmt"
)

/*
	给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
	你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
	返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

	输入：[7,1,5,3,6,4]
	输出：5
	解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
		 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

	链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

*/

func main() {
	price := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(price))
}

func maxProfit(prices []int) int {
	var dp = make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < len(prices); i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			// 解释：
			//   dp[0][0]
			// = max(dp[-1][0], dp[-1][1] + prices[i])
			// = max(0, -infinity + prices[i]) = 0
			dp[i][1] = -prices[i]
			//解释：
			//   dp[0][1]
			// = max(dp[-1][1], dp[-1][0] - prices[i])
			// = max(-infinity, 0 - prices[i])
			// = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	fmt.Println(dp)
	return dp[len(prices)-1][0]
}

//{7, 1, 5, 3, 6, 4}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
