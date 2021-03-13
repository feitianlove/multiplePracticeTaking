package main

/*
	给你一个字符串 s 和一个整数 k 。请你用 s 字符串中 所有字符 构造 k 个非空 回文串 。
	如果你可以用 s 中所有字符构造 k 个回文字符串，那么请你返回 True ，否则返回 False 。

	链接：https://leetcode-cn.com/problems/construct-k-palindrome-strings
	输入：s = "annabelle", k = 2

	输出：true
	解释：可以用 s 中所有字符构造 2 个回文字符串。
	一些可行的构造方案包括："anna" + "elble"，"anbna" + "elle"，"anellena" + "b"

*/
func main() {
	canConstruct("annabelle", 3)
}

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	// a = 0, b = 1
	//var arr = make([]int,27)
	//for i := 0; i < len(s); i++ {
	//	arr[s[i]-'a']++
	//}
	//a = 97, b = 98
	var arr = make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		arr[s[i]]++
	}
	res := 0
	for _, value := range arr {
		if value%2 == 1 {
			res++
		}
	}
	return res <= k
}
