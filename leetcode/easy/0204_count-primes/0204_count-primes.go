package main

import "fmt"

/*
	统计所有小于非负整数 n 的质数的数量。

	输入：n = 10
	输出：4
	解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

	链接：https://leetcode-cn.com/problems/count-primes

*/

func main() {
	fmt.Println(countPrimes(10))
}

// 埃氏筛法
func countPrimes(n int) int {
	var arr = make([]bool, n)
	var count int
	for i := 2; i*i < n; i++ {
		if !arr[i] {
			for j := i * i; j < n; j += i {
				arr[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !arr[i] {
			count++
		}
	}
	return count
}
