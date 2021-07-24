package main

import "fmt"

/*
	给定两个字符串s1, s2，找到使两个字符串相等所需删除字符的ASCII值的最小和。

	输入: s1 = "sea", s2 = "eat"
	输出: 231
	解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
	在 "eat" 中删除 "t" 并将 116 加入总和。
	结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
	示例 2:

	输入: s1 = "delete", s2 = "leet"
	输出: 403
	解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
	将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
	结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
	如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。

	链接：https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings
*/
//"delete"
//"delete"
func main() {
	fmt.Println(minimumDeleteSum("delete", "leet"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1)
		if i == 0 {
			continue
		}
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for j := 0; j <= len(s2); j++ {
		if j == 0 {
			continue
		}

		dp[0][j] += dp[0][j-1] + int(s2[j-1])
	}

	//"delete", "let"
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				s1Value := int(s1[i-1])
				s2Value := int(s2[j-1])
				//e  ea
				dp[i][j] = min(dp[i-1][j]+s1Value, dp[i][j-1]+s2Value)
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
