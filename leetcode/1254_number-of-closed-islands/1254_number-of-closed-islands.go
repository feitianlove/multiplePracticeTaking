package main

import "fmt"

/*

	有一个二维矩阵 grid ，每个位置要么是陆地（记号为 0 ）要么是水域（记号为 1 ）。
	我们从一块陆地出发，每次可以往上下左右 4 个方向相邻区域走，能走到的所有陆地区域，我们将其称为一座「岛屿」。
	如果一座岛屿 完全 由水域包围，即陆地边缘上下左右所有相邻区域都是水域，那么我们将其称为 「封闭岛屿」。
	请返回封闭岛屿的数目。

	输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
	输出：2
	解释：
	灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。

	链接：https://leetcode-cn.com/problems/number-of-closed-islands

*/

func main() {
	arr := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	fmt.Println(closedIsland(arr))
}
func closedIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				if dfs(grid, i, j) == true {
					res++
				}
			}
		}
	}
	return res
}
func dfs(grid [][]int, i, j int) bool {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}
	if grid[i][j] == 1 {
		return true
	}
	grid[i][j] = 1
	// 下面写法有问题，与逻辑执行如果返回false后面不执行了
	//return dfs(grid, i, j+1) && dfs(grid, i, j-1) && dfs(grid, i+1, j) && dfs(grid, i+1, j)
	up := dfs(grid, i, j+1)
	down := dfs(grid, i, j-1)
	left := dfs(grid, i-1, j)
	right := dfs(grid, i+1, j)
	return up && down && left && right
}
