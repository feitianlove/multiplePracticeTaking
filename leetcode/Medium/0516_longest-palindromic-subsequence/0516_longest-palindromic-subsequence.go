package main

import "fmt"

/*
	给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
	子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

	输入：s = "bbbab"
	输出：4
	解释：一个可能的最长回文子序列为 "bbbb" 。


	链接：https://leetcode-cn.com/problems/longest-palindromic-subsequence
*/
func main() {
	fmt.Println(longestPalindromeSubseq("b"))
}
func longestPalindromeSubseq(s string) int {
	l := len(s)
	var dp = make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		dp[i][i] = 1
	}
	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][l-1]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
