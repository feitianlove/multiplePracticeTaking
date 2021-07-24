package main

import "fmt"

/*
	给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

	输入: "sea", "eat"
	输出: 2
	解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"

	链接：https://leetcode-cn.com/problems/delete-operation-for-two-strings
*/
func main() {
	fmt.Println(minDistance("sea", "eat"))
}

//"sea"
//"eat"
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
