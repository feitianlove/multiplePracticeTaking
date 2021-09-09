package main

import "fmt"

/*
	数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
	有效括号组合需满足：左括号必须以正确的顺序闭合。


	输入：n = 3
	输出：["((()))","(()())","(())()","()(())","()()()"]

	链接：https://leetcode-cn.com/problems/generate-parentheses
*/
func main() {
	fmt.Println(generateParenthesis(3))
}

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	backtrack(n, n, "")
	return res

}

func backtrack(left, right int, track string) {
	if left < 0 || right < 0 {
		return
	}
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		res = append(res, track)
		return
	}
	track += "("
	backtrack(left-1, right, track)
	track = track[:len(track)-1]

	track += ")"
	backtrack(left, right-1, track)
	track = track[:len(track)-1]
}
