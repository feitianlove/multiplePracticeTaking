package main

import "fmt"

/*
	给定一个包含了一些 0 和 1 的非空二维数组 grid 。

	一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

	找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

	[[0,0,1,0,0,0,0,1,0,0,0,0,0],
	 [0,0,0,0,0,0,0,1,1,1,0,0,0],
	 [0,1,1,0,1,0,0,0,0,0,0,0,0],
	 [0,1,0,0,1,1,0,0,1,0,1,0,0],
	 [0,1,0,0,1,1,0,0,1,1,1,0,0],
	 [0,0,0,0,0,0,0,0,0,0,1,0,0],
	 [0,0,0,0,0,0,0,1,1,1,0,0,0],
	 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
	对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。

	链接：https://leetcode-cn.com/problems/max-area-of-island
*/
func main() {
	arr := [][]int{{0, 1, 0, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(arr))
}
func maxAreaOfIsland(grid [][]int) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				temp := dfs(grid, i, j)
				if temp > res {
					res = temp
				}
			}
		}
	}

	return res
}
func dfs(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || j >= len(grid[0]) || i >= len(grid) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	res := 1
	res = res + dfs(grid, i+1, j)
	res = res + dfs(grid, i-1, j)
	res = res + dfs(grid, i, j+1)
	res = res + dfs(grid, i, j-1)
	return res
}
