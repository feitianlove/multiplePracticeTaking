package main

import (
	"fmt"
	"math"
)

/*
	给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
	你可以对一个单词进行如下三种操作：

	插入一个字符
	删除一个字符
	替换一个字符


	输入：word1 = "horse", word2 = "ros"
	输出：3
	解释：
	horse -> rorse (将 'h' 替换为 'r')
	rorse -> rose (删除 'r')
	rose -> ros (删除 'e')

	链接：https://leetcode-cn.com/problems/edit-distance
*/
func main() {
	fmt.Println(minDistance2("horse", "ros"))
}

// 递归解法+备忘录
var memo [][]int

func minDistance(word1 string, word2 string) int {
	memo = make([][]int, len(word1))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(word2))
	}
	return dp(word1, word2, len(word1)-1, len(word2)-1)
}

func dp(word1 string, word2 string, i, j int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	var res int
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	if word1[i] == word2[j] {
		res = dp(word1, word2, i-1, j-1)
	} else {
		//增加、删除、替换
		res = min(dp(word1, word2, i, j-1)+1,
			dp(word1, word2, i-1, j)+1,
			dp(word1, word2, i-1, j-1)+1)
	}
	memo[i][j] = res
	return res
}

//动态规划
func minDistance2(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	w1 := len(word1)
	w2 := len(word2)
	// dp是长度从0-i的word1到0-j的word2的最小移动次数
	for i := 0; i <= w1; i++ {
		dp[i] = make([]int, w2+1)
		dp[i][0] = i
	}
	// a b c
	// d e f
	for i := 0; i < w2+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= w1; i++ {
		for j := 1; j <= w2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j-1]+1, dp[i-1][j]+1)
			}
		}
	}
	return dp[w1][w2]
}

func min(a, b, c int) int {
	a_a := float64(a)
	b_b := float64(b)
	c_c := float64(c)
	return int(math.Min(a_a, math.Min(b_b, c_c)))
}
