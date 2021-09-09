package main

import "fmt"

/*
	给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
	你可以按 任何顺序 返回答案。

	输入：n = 4, k = 2
	输出：
	[
	  [2,4],
	  [3,4],
	  [2,3],
	  [1,2],
	  [1,3],
	  [1,4],
	]

	链接：https://leetcode-cn.com/problems/combinations
*/
func main() {
	combine(4, 2)
	fmt.Println(result)
}

var result [][]int

func combine(n int, k int) [][]int {
	result = make([][]int, 0)
	var track = make([]int, 0)
	backtrack(n, 1, k, track)
	return result
}

func backtrack(n, start, k int, track []int) {
	if k == len(track) {
		var temp = make([]int, len(track))
		copy(temp, track)
		result = append(result, temp)
		return
	}
	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(n, i+1, k, track)
		track = track[:len(track)-1]
	}
}
