package main

import (
	"fmt"
)

/*
	给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
	输入：n = 3
	输出：5

	https://leetcode-cn.com/problems/unique-binary-search-trees/
*/
func main() {
	fmt.Println(numTrees(5))
}

var dp [][]int

func numTrees(n int) int {
	dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	return count(1, n)
}

func count(start, end int) int {
	var res int

	if start >= end {
		return 1
	}
	if dp[start][end] != 0 {
		//fmt.Println(start, end)
		return dp[start][end]
	}
	//[1,2,3,4,5]
	for i := start; i <= end; i++ {
		left := count(start, i-1)
		right := count(i+1, end)
		res += left * right
	}
	dp[start][end] = res
	return res
}
