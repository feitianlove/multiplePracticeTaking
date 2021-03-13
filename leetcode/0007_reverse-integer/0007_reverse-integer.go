package main

/*
	给你一个 32 位的有符号整数 x ，返回 x 中每位上的数字反转后的结果。
	如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
	假设环境不允许存储 64 位整数（有符号或无符号）。

	链接：https://leetcode-cn.com/problems/reverse-integer
*/
import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	var result string
	str := strconv.Itoa(x)
	var r []string
	r = strings.Split(str, "")
	if x < 0 {
		r = r[1:]
	}

	for i := len(r) - 1; i >= 0; i-- {
		result += r[i]
	}
	re, _ := strconv.Atoi(result)
	if re > math.MaxInt32 || re < math.MinInt32 {
		re = 0
	}
	if x < 0 {
		return -re
	}
	return re
}
func reverse2(x int) int {
	var result string
	flag := 1
	if x < 0 {
		flag = -1
		x = x * -1
	}
	for x > 0 {
		temp := x % 10
		x = x / 10
		result += strconv.Itoa(temp)
	}
	r, _ := strconv.Atoi(result)
	if r > math.MaxInt32 {
		r = 0
	}
	return flag * r
}
func reverse3(x int) int {
	result := 0
	for x != 0 {
		temp := x % 10
		fmt.Println(temp, x)
		result = result*10 + temp
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return result
}
func main() {
	result := reverse3(-123)
	fmt.Println(result)
}
