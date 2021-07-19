package main

/*

	给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
	下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列
	（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、
	(row + 1, col) 或者 (row + 1, col + 1) 。

	输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
	输出：13
	解释：下面是两条和最小的下降路径，用加粗标注：
	[[2,1,3],      [[2,1,3],
	 [6,5,4],       [6,5,4],
	 [7,8,9]]       [7,8,9]]

	链接：https://leetcode-cn.com/problems/minimum-falling-path-sum

*/

import (
	"fmt"
	"math"
)

func main() {
	var matrix [][]int = [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(matrix))
}

var dp [][]int

func minFallingPathSum(matrix [][]int) int {
	dp = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	var res = math.MaxInt32
	for j := 0; j < len(matrix[0]); j++ {
		res = min(res, dfs(matrix, 0, j))
	}
	return res
}
func dfs(matrix [][]int, i, j int) int {

	if j < 0 || j >= len(matrix[i]) {
		return math.MaxInt32
	}
	if i == len(matrix)-1 {
		return matrix[i][j]
	}
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	res := matrix[i][j] + minThreee(dfs(matrix, i+1, j),
		dfs(matrix, i+1, j-1),
		dfs(matrix, i+1, j+1))
	dp[i][j] = res
	return res
}

func minThreee(a, b, c int) int {
	a_a := float64(a)
	b_b := float64(b)
	c_c := float64(c)
	return int(math.Min(a_a, math.Min(b_b, c_c)))
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
