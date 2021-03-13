package main

import "fmt"

/*
	给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
	你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
	返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

	链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

*/

func main() {
	var arr = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(arr))
}

// 超出时间限制
func maxProfit(prices []int) int {
	max := -1
	length := len(prices)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	if max > 0 {
		return max
	} else {
		return 0
	}
}
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	profit := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if max < prices[i] {
			max = prices[i]
		}
		if profit < max-prices[i] {
			profit = max - prices[i]
		}
	}
	return profit
}
