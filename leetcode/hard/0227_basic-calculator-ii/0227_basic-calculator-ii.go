package main

import (
	"fmt"
	"strings"
)

/*
	给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
	整数除法仅保留整数部分。



	输入：s = "3+2*2"
	输出：7


	链接：https://leetcode-cn.com/problems/basic-calculator-ii
*/

func main() {

	s := "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(calculate(s))

}

func calculate(s string) int {
	return cc(&s)
}

func cc(ss *string) int {
	s := *ss
	s = strings.Trim(s, " ")
	var stack = make([]int, 0)
	var sign = '+'
	var num = 0
	for len(s) > 0 {
		ch := s[0]
		s = s[1:]
		if '0' <= ch && ch <= '9' {
			num = num*10 + int(ch-'0')
		}
		if ch == '(' {
			num = cc(&s)
		}
		if ch != ' ' && ('0' > ch || ch > '9' || 0 == len(s)) {
			var pre int
			switch sign {
			case '+':
				stack = append(stack, +num)
			case '-':
				stack = append(stack, -num)
			case '*':
				pre = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pre*num)
			case '/':
				pre = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pre/num)
			}
			num = 0
			sign = int32(ch)
		}
		if ch == ')' {
			*ss = s
			break
		}
	}
	var res = 0
	for i := len(stack) - 1; i >= 0; i-- {
		res += stack[i]
	}
	return res
}
