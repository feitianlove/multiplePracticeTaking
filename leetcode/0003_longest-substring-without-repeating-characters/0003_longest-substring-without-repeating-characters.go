package main

/*

	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
	输入: s = "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
	https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	start := 0
	for i := 1; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			continue
		}

		if i-start > max {
			max = i - start
		}
		start = start + index + 1
	}
	if len(s)-start > max {
		max = len(s) - start
	}
	return max
}
