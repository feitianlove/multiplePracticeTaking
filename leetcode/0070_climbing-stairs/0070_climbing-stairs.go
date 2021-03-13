package main

import "fmt"

func main() {
	fmt.Println(climbStairs2(6))
}

//用数组优化
var arr []int

// 递归
func climbStairs(n int) int {
	arr = make([]int, n+1)
	return start(n)

}
func start(n int) int {
	if n == 1 {
		arr[n] = 1
		return 1
	}
	if n == 2 {
		arr[n] = 2
		return 2
	}
	if arr[n] > 0 {
		return arr[n]
	}
	arr[n] = start(n-1) + start(n-2)
	return arr[n]
}

// 动态规划
func climbStairs2(n int) int {
	var dp = make([]int, n+1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
