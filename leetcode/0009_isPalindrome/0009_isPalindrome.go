package main

/*
	给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
	回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

	链接：https://leetcode-cn.com/problems/palindrome-number
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := isPalindrome2(121)
	fmt.Println(result)
}

//1 2 3 4 5
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	r := strings.Split(str, "")
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-i-1] {
			return false
		}
	}
	return true
}
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	result := 0
	// x = 1221  => x = 12 revertedNumber = 12
	// x = 12321 => x = 12 revertedNumber = 123
	for x > result {
		temp := x % 10
		result = result*10 + temp
		x = x / 10
	}
	return result == x || result/10 == x
}
