package main

import (
	"fmt"
	"math"
	"strings"
)

/*
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
	注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

	输入：s = "ADOBECODEBANC", t = "ABC"
	输出："BANC"

	链接：https://leetcode-cn.com/problems/minimum-window-substring

*/

func main() {
	s := "bbaa"
	t := "aba"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	left, right := 0, 0
	// 记录窗口中的字符
	word := make(map[uint8]int)
	//记录需要的字符
	need := make(map[uint8]int)
	valid := 0
	dd := ""
	var res = math.MaxInt32
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	for right < len(s) {
		c := s[right]
		right++
		if len(need) > valid {
			if strings.Contains(t, string(c)) {
				word[c]++
				if word[c] == need[c] {
					valid++
				}
			}
		}

		for valid == len(need) {
			if right-left < res {
				res = right - left
				dd = s[left:right]
			}
			c2 := s[left]
			left++
			if strings.Contains(t, string(c2)) {
				if word[c2] == need[c2] {
					valid--
				}
				word[c2]--
			}
		}
	}
	return dd
}
