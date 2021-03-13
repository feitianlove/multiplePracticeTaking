package main

import (
	"fmt"
)

/*
	给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
	https://leetcode-cn.com/problems/pascals-triangle-ii/
*/

func main() {
	fmt.Println(getRow(2))
}

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}
	for i := 1; i <= rowIndex; i++ {
		for j := i; j-1 >= 0; j-- {
			res[j] = res[j] + res[j-1]
		}
	}
	return res
}
