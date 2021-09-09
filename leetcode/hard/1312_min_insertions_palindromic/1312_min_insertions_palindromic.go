package main

import "fmt"

/*
	给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
	请你返回让 s 成为回文串的 最少操作次数 。
	「回文串」是正读和反读都相同的字符串。

	输入：s = "zzazz"
	输出：0
	解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。

	链接：https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome
*/

func main() {
	fmt.Println(minInsertions("za"))
}

func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	//bash case  dp[i][j] 当 i=j  ,dp[i][j]= 0, 不等于的时候如下
	//转移方程
	//1、当s[i] == s[j]:  dp[i][j] =
	//2、当s[i] != s[j]:  要排除  aaaaab 不能直接 dp[i+1][j-1]+2
	//		先计算dp[i+1][j] 或者dp[i][j-1],取最小的
	//  min(dp[i+1][j] , dp[i][j-1]) +1
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
