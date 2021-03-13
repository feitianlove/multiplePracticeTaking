package main

/*
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。
	输入：strs = ["flower","flow","flight"]
	输出："fl"
*/

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	strs := []string{"flower", "flow", "flight"}
	//result := longestCommonPrefix2(strs)
	//result := longestCommonPrefix(strs)
	result := longestCommonPrefix3(strs)

	fmt.Println(result)
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j]) || char != strs[j][i] {
				return strs[0][:i]
			}
		}

	}
	return strs[0]
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	short := strs[0]
	for _, item := range strs {
		if len(short) > len(item) {
			short = item
		}
	}
	for i := range short {
		shortest := short[:i+1]
		for _, str := range strs {
			if strings.Index(str, shortest) != 0 {
				return short[:i]
			}
		}
	}
	return short
}

func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)
	// TODO  go中slice的排序
	//sort.SliceStable(strs, func(i, j int) bool {
	//	return len(strs[i]) > len(strs[j])
	//})
	first := strs[0]
	last := strs[len(strs)-1]
	i := 0
	length := len(first)
	if len(last) < length {
		length = len(last)
	}
	for i < length {
		if first[i] != last[i] {
			return first[:i]
		}
		i++
	}
	return first[:i]
}

// TODO 4，5，6 三种方法没写
