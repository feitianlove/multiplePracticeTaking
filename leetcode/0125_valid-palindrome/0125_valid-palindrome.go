package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
	说明：本题中，我们将空字符串定义为有效的回文串。
	输入: "A man, a plan, a canal: Panama"
	输出: true
	链接：https://leetcode-cn.com/problems/valid-palindrome
*/

func main() {
	fmt.Println(isPalindrome("race a car"))
}
func isPalindrome(s string) bool {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	result := reg.ReplaceAllString(s, "")
	result = strings.ToUpper(result)
	middle := len(result) / 2
	length := len(result) - 1
	for i := 0; i < middle; i++ {
		if result[i] != result[length-i] {
			return false
		}
	}
	return true
}
