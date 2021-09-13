package main

import (
	"fmt"
	"math"
)

/*
	给定一个整数 n ，返回 n! 结果中尾随零的数量。
	进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？

	输入：n = 3
	输出：0
	解释：3! = 6 ，不含尾随 0

	链接：https://leetcode-cn.com/problems/factorial-trailing-zeroes
*/
func main() {
	fmt.Println(trailingZeroes(125))
	fmt.Println(trailingZeroes(math.MaxInt32))
}

/*
	也就是找2*5的因子，5的因子明显较少，所以 125/5 = 25个
	但是25的倍数的可以额外提供一个5*5
	125的也可以额外提供一个5*5*5
*/
func trailingZeroes(n int) int {
	var res int
	div := 5
	for n >= div {
		res += n / div
		div *= 5
	}
	return res
}
