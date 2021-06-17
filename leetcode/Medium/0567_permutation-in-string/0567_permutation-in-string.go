package main

import (
	"fmt"
	"strings"
)

/*
	给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
	换句话说，第一个字符串的排列之一是第二个字符串的 子串 。

	输入: s1 = "ab" s2 = "eidbaooo"
	输出: True
	解释: s2 包含 s1 的排列之一 ("ba")

	https://leetcode-cn.com/problems/permutation-in-string/
*/

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[uint8]int)
	word := make(map[uint8]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	vaild := 0
	left, right := 0, 0
	for right < len(s2) {
		c := s2[right]
		right++
		if vaild < len(need) {
			if strings.Contains(s1, string(c)) {
				word[c]++
				if word[c] == need[c] {
					vaild++
				}
			}
		}
		for right-left >= len(s1) {
			if vaild == len(need) {
				return true
			}
			c := s2[left]
			left++
			if strings.Contains(s1, string(c)) {
				if word[c] == need[c] {
					vaild--
				}
				word[c]--
			}
		}
	}
	return false
}
