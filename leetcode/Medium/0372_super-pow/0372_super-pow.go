package main

import "fmt"

/*
	你的任务是计算 ab 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。

	输入：a = 2, b = [3]
	输出：8


	链接：https://leetcode-cn.com/problems/super-pow
*/
func main() {
	fmt.Println(superPow(2, []int{4}))
}

//(a % k)(b % k) % k = BD % k
var base = 1337

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]

	part1 := myPow(a, last)
	part2 := myPow(superPow(a, b), 10)

	return (part1 * part2) % base
}

//
func myPow1(a, k int) int {
	a %= base
	var res = 1
	for i := 1; i <= k; i++ {
		res *= a
		res %= base
	}
	return res
}

// 优化的myPow

func myPow(a, k int) int {
	if k == 0 {
		return 1
	}
	a %= base
	if k%2 == 1 {
		return a * myPow(a, k-1)
	} else {
		sub := myPow(a, k/2)
		return (sub * sub) % base
	}
}
