package main

import "fmt"

/*
	给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

	说明：
	拆分时可以重复使用字典中的单词。
	你可以假设字典中没有重复的单词。

	示例 1：
	输入: s = "leetcode", wordDict = ["leet", "code"]
	输出: true
	解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

	链接：https://leetcode-cn.com/problems/word-break

*/
func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	for i := 1; i <= len(s); i++ {
		dp[0] = true
		for j := 0; j < i; j++ {
			if dp[j] == true && m[s[j:i]] == true {
				dp[i] = true
			}

		}
	}
	return dp[len(s)]
}
