package main

import (
	"fmt"
	"strconv"
)

/*
	给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。

	输入: "2-1-1"
	输出: [0, 2]
	解释:
	((2-1)-1) = 0
	(2-(1-1)) = 2

	链接：https://leetcode-cn.com/problems/different-ways-to-add-parentheses
*/
func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))

}

func diffWaysToCompute(expression string) []int {
	var res = make([]int, 0)
	for i := 0; i < len(expression); i++ {
		char := expression[i]
		if char == '+' || char == '-' || char == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch char {
					case '-':
						res = append(res, l-r)
					case '+':
						res = append(res, l+r)
					case '*':
						res = append(res, l*r)

					}
				}
			}
		}
	}
	if len(res) == 0 {
		temp, _ := strconv.Atoi(expression)
		return []int{temp}
	}
	return res
}
