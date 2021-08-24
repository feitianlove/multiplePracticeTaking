package main

import (
	"fmt"
	"math"
)

/*
	有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城市 fromi 开始，以价格
	toi 抵达 pricei。现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst
	的 价格最便宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。

	输入:
	n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
	src = 0, dst = 2, k = 1
	输出: 200
	从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200

	链接：https://leetcode-cn.com/problems/cheapest-flights-within-k-stops
*/
func main() {
	n := 3
	src := 0
	dst := 2
	k := 1
	flights := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))
}

var indegree map[int][][]int
var memo [][]int

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	indegree = make(map[int][][]int)
	memo = make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, k+2)
	}
	for _, item := range flights {
		from := item[0]
		to := item[1]
		price := item[2]
		indegree[to] = append(indegree[to], []int{from, price})
	}
	return dp(src, dst, k+1)
}

func dp(src, des int, k int) int {
	if src == des {
		return 0
	}
	if k == 0 {
		return -1
	}
	res := math.MaxInt32

	if memo[des][k] != 0 {
		return memo[des][k]
	}
	data, ok := indegree[des]
	if ok {
		for _, item := range data {
			from := item[0]
			price := item[1]
			subProblem := dp(src, from, k-1)
			if subProblem != -1 {
				res = min(res, subProblem+price)
			}
		}
	}
	if res == math.MaxInt32 {
		memo[des][k] = -1
	} else {
		memo[des][k] = res
	}

	return memo[des][k]
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
