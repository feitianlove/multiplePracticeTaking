package main

import "fmt"

/*
	给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
	说明：每次只能向下或者向右移动一步。

	输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
	输出：7
	解释：因为路径 1→3→1→1→1 的总和最小。

	链接：https://leetcode-cn.com/problems/minimum-path-sum
*/
func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	fmt.Println(minPathSum(grid))
}

// 1 2
// 3 4
func minPathSum(grid [][]int) int {
	var dp [][]int = make([][]int, len(grid))
	var temp int
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
		temp += grid[i][0]
		dp[i][0] = temp
	}
	temp = 0
	for i := 0; i < len(grid[0]); i++ {
		temp += grid[0][i]
		dp[0][i] = temp
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
