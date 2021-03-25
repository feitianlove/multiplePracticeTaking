package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
	链接：https://leetcode-cn.com/problems/string-to-integer-atoi

*/
func main() {
	fmt.Println(myAtoi("42"))
}

func myAtoi(s string) int {
	st := strings.Trim(s, " ")
	if len(st) == 0 {
		return 0
	}
	res := ""
	flag := 1
	if st[0] == '-' {
		flag = -1
		st = st[1:]
	} else if st[0] == '+' {
		flag = 1
		st = st[1:]
	}
	for i := 0; i < len(st); i++ {

		if '0' <= st[i] && st[i] <= '9' {
			res += string(st[i])
		} else {
			break
		}
	}
	n, _ := strconv.Atoi(res)
	if n*flag > math.MaxInt32 {
		return math.MaxInt32
	} else if n*flag < math.MinInt32 {
		return math.MinInt32
	} else {
		return n * flag
	}
}
