package main

import (
	"fmt"
	"strconv"
)

/*
	给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
	所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。


	示例 1：

	输入：s = "aa" p = "a"
	输出：false
	解释："a" 无法匹配 "aa" 整个字符串。
	示例 2:

	输入：s = "aa" p = "a*"
	输出：true
	解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
	示例 3：

	输入：s = "ab" p = ".*"
	输出：true
	解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
	示例 4：

	输入：s = "aab" p = "c*a*b"
	输出：true
	解释：因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

	链接：https://leetcode-cn.com/problems/regular-expression-matching
*/

func main() {
	//fmt.Println(isMatch("a", "ab*c*"))
	fmt.Println(isMatch("ab", ".*"))

}

var memo map[string]bool

func isMatch(s string, p string) bool {
	memo = make(map[string]bool)
	return dp(s, 0, p, 0)
}

func dp(s string, i int, p string, j int) bool {
	if len(p) == j {
		return i == len(s)
	}
	// a ab*c*
	if i == len(s) {
		if (len(p)-j)%2 == 1 {
			return false
		}
		for ; j < len(p)-1; j = j + 2 {
			if p[j+1] != '*' {
				return false
			}
		}
		return true
	}
	key := strconv.Itoa(i) + "-" + strconv.Itoa(j)
	if value, ok := memo[key]; ok {
		return value
	}
	res := false
	if s[i] == p[j] || p[j] == '.' {
		if j < len(p)-1 && p[j+1] == '*' {
			//abcd  aba*cd 匹配0次 ||abaaac   aba*c
			//return dp(s, i, p, j+2) || dp(s, i+1, p, j)
			res = dp(s, i, p, j+2) || dp(s, i+1, p, j)
		} else {
			res = dp(s, i+1, p, j+1)
		}
	} else {
		if j < len(p)-1 && p[j+1] == '*' {
			//abcd  aba*cd 匹配0次
			res = dp(s, i, p, j+2)
		} else {
			res = false
		}
	}
	memo[key] = res
	return res
}
