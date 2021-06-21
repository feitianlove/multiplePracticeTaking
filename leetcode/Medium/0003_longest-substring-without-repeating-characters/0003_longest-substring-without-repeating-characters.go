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
	//fmt.Println(lengthOfLongestSubstring2("abc"))
	//fmt.Println(isRepeat("abca"))
	fmt.Println(lengthOfLongestSubstring3("tmmzuxt"))
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

// 暴力法
func lengthOfLongestSubstring2(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isRepeat(s[i:j]) {
				if len(s[i:j]) > max {
					max = len(s[i:j])
				}
			}
		}
	}
	return 0
}
func isRepeat(s string) bool {
	for i := 0; i < len(s); i++ {
		index := strings.Index(s, string(s[i]))
		if index != -1 {
			return false
		}
	}
	return true
}

//滑动窗口
func lengthOfLongestSubstring3(s string) int {
	//tmmzuxt  abcabcbb
	var left, right = 0, 0
	var res int
	words := make(map[uint8]int)
	for right < len(s) {
		s1 := s[right]
		right++
		words[s1]++
		for words[s1] > 1 {
			s2 := s[left]
			left++
			words[s2]--

		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
