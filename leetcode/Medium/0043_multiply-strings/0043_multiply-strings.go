package main

import (
	"fmt"
	"strconv"
)

/*
	给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

	输入: num1 = "123", num2 = "456"
	输出: "56088"

	链接：https://leetcode-cn.com/problems/multiply-strings
*/

func main() {
	fmt.Println(multiply("456", "123"))
	//fmt.Println(multiply("10", "2"))

}

func multiply(num1 string, num2 string) string {
	m := len(num1)
	n := len(num2)
	var res = make([]uint8, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			tempI := num1[i] - '0'
			tempJ := num2[j] - '0'
			r := tempJ*tempI + res[i+j+1]
			res[i+j] += r / 10
			res[i+j+1] = r % 10
		}
	}
	var i = 0
	var result string
	for i < len(res) && res[i] == 0 {
		i++
	}
	for j := i; j < len(res); j++ {
		result += strconv.Itoa(int(res[j]))
	}
	if len(res) == 0 {
		return "0"
	}
	return result
}
