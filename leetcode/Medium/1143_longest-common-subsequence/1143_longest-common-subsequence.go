package main

import "fmt"

/*
	给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
	一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）
	后组成的新字符串。

	例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
	两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

	输入：text1 = "abcde", text2 = "ace"
	输出：3
	解释：最长公共子序列是 "ace" ，它的长度为 3

	链接：https://leetcode-cn.com/problems/longest-common-subsequence

*/

func main() {
	fmt.Println(longestCommonSubsequence1("mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq"))
}

var dp [][]int

func longestCommonSubsequence(text1 string, text2 string) int {
	dp = make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
	}
	return dfs(text1, 0, text2, 0)
}

// 递归+备忘录
func dfs(s1 string, i int, s2 string, j int) int {
	if i == len(s1) || j == len(s2) {
		return 0
	}
	if dp[i][j] != 0 {
		fmt.Println(dp[i][j])
		return dp[i][j]
	}
	if s1[i] == s2[j] {
		dp[i][j] = 1 + dfs(s1, i+1, s2, j+1)
	} else {

		dp[i][j] = max(
			dfs(s1, i+1, s2, j),
			dfs(s1, i, s2, j+1),
			//dfs(s1, i+1, s2, j+1), 这个比dfs(s1, i, s2, j+1)短，所以可以舍弃
		)
	}
	return dp[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//动态规划
//"ab", text2 = "ac"
func longestCommonSubsequence1(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
