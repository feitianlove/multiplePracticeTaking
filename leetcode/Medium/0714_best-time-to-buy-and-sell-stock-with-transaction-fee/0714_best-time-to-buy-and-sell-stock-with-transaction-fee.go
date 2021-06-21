package main

import (
	"fmt"
	"math"
)

/*
	给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
	你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
	返回获得利润的最大值。

	注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费

	输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
	输出: 8
	解释: 能够达到的最大利润:
	在此处买入 prices[0] = 1
	在此处卖出 prices[3] = 8
	在此处买入 prices[4] = 4
	在此处卖出 prices[5] = 9
	总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.

	链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee

*/

func main() {
	price := []int{1, 3, 2, 8, 4, 9}

	fmt.Println(maxProfit(price, 2))
}

func maxProfit(prices []int, fee int) int {
	var n = len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt32
	pre_dp := 0
	for i := 0; i < n; i++ {
		pre_dp = dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i]-fee)
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
