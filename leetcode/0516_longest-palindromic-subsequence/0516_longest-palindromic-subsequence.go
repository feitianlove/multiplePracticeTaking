package main

import "fmt"

/*
	给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。
	"bbbab"
	输出: 4
	一个可能的最长回文子序列为 "bbbb"。

	https://leetcode-cn.com/problems/longest-palindromic-subsequence/
*/

func main() {
	fmt.Println(longestPalindromeSubseq("bbab"))
}

func longestPalindromeSubseq(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	n := len(s)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
