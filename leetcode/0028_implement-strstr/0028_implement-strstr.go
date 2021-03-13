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
	haystack := "0"
	needle := "0"
	ret := strStr2(haystack, needle)
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

func strStr3(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	return 0
}

// TODO sunday 算法
func strStr4(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	return 0
}
