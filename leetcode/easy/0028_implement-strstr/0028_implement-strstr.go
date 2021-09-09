package main

import (
	"fmt"
	"strings"
)

/*
	strstr(str1,str2) 函数用于判断字符串str2是否是str1的子串。如果是，则该函数返回 str1字符串从 str2第一次出现的位置开始到 str1结
	尾的字符串；否则，返回NULL。
*/
func main() {
	haystack := ""
	needle := ""
	//ret := strStr2(haystack, needle)
	ret := kmp(haystack, needle)
	fmt.Println(ret)
}

// 相当于实现strings.Index方法
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	ret := strings.Index(haystack, needle)
	return ret
}

func strStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:nlen+i] == needle {
			return i
		}
	}
	return -1
}

// TODO KMP 算法 https://wiki.jikexueyuan.com/project/kmp-algorithm/define.html
func kmp(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	var dp = make([][256]int, n)
	// base case 只有遇到第一个字母才能从0->1
	dp[0][needle[0]] = 1
	// 影子状态初始化为0
	var X int
	for i := 1; i < n; i++ {
		for c := 0; c < 256; c++ {
			if int(needle[i]) == c {
				dp[i][c] = i + 1
			} else {
				dp[i][c] = dp[X][c]
			}
		}
		//abab ->c如果遇到 b 则0 ，如果遇到 a则 会推到2
		// 更新影子状态
		X = dp[X][needle[i]]
	}

	m := len(haystack)
	j := 0
	for i := 0; i < m; i++ {
		j = dp[j][haystack[i]]
		if j == n {
			return i - n + 1
		}
	}
	return -1
}
