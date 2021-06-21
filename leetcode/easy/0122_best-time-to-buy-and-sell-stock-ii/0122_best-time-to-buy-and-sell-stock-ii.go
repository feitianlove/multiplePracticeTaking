package main

import (
	"fmt"
	"math"
)

/*

	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
	设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

	链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
	输入: prices = [7,1,5,3,6,4]
	输出: 7
	解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

*/

func main() {
	var prices = []int{7, 1, 3, 5, 6, 4}
	//
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit1(prices))

}

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < len(prices); i++ {
		if i-1 == -1 {
			//dp[0][0]
			//= max (dp[-1][0], dp[-1][1]+price)
			//= max(0, -infinity + prices[i])
			dp[i][0] = 0
			// dp[0][1]
			//= max (dp[-1][1], -price)
			//= max(-infinity , - prices[i])
			dp[i][1] = -prices[i]
			continue
		}
		//7, 1, 5, 3, 6, 4
		//7  [0 -7]
		//1  [0, -1]
		//5  [4, -1]
		// 第二次
		//3  [4 1]
		//6   [7 1]

		//7, 1, 3, 5, 6, 4
		//7  [0 -7]
		//1  [0 -1]
		//3  [2 -1]
		//5  [4 -1]
		//6  [5 1]
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	fmt.Println(dp)
	return dp[len(prices)-1][0]
}

func maxProfit1(prices []int) int {
	var n = len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt32
	var pre_dp = 0
	for i := 0; i < n; i++ {
		pre_dp = dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, pre_dp-prices[i])
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
