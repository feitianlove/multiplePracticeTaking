package main

import (
	"fmt"
	"strings"
)

/*
	给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
	字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
	说明：
		字母异位词指字母相同，但排列不同的字符串。
		不考虑答案输出的顺序
	输入:
	s: "abab" p: "ab"
	输出:
	[0, 1, 2]
	解释:
	起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
	起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
	起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。


	链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string

*/

func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	vaild := 0
	need := make(map[uint8]int)
	words := make(map[uint8]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	var left, right = 0, 0
	for right < len(s) {
		s1 := s[right]
		right++
		if vaild < len(need) {
			if strings.Contains(p, string(s1)) {
				words[s1]++
				if words[s1] == need[s1] {
					vaild++
				}
			}
		}
		if right-left >= len(p) {
			s2 := s[left]
			if len(need) == vaild {
				res = append(res, left)
			}
			left++
			if strings.Contains(p, string(s2)) {
				if words[s2] == need[s2] {
					vaild--
				}
				words[s2]--
			}
		}
	}

	return res
}
