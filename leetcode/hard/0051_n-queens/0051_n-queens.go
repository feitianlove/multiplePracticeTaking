package main

import "fmt"

/*
	n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
	给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
	每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

	链接：https://leetcode-cn.com/problems/n-queens


	输入：n = 4
	输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
	解释：如上图所示，4 皇后问题存在两个不同的解法。

*/
func main() {
	fmt.Println(solveNQueens(4))
}

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	arr := make([][]string, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]string, n)
		for j := 0; j < n; j++ {
			arr[i][j] = "."
		}
	}
	dfs(arr, 0)
	return res
}

func dfs(arr [][]string, row int) {
	if len(arr) == row {
		temp := make([]string, 0)
		for i := 0; i < row; i++ {
			st := ""
			for j := 0; j < len(arr[i]); j++ {
				st += arr[i][j]
			}
			temp = append(temp, st)
		}
		res = append(res, temp)
	}

	for col := 0; col < len(arr[0]); col++ {
		if !isValid(arr, row, col) {
			continue
		}
		arr[row][col] = "Q"
		dfs(arr, row+1)
		arr[row][col] = "."
	}
}

func isValid(arr [][]string, row, col int) bool {
	//先判断同列有没有
	n := len(arr)
	// 当前列判断(竖着)
	for row := 0; row < n; row++ {
		if arr[row][col] == "Q" {
			return false
		}
	}
	//判断左上角
	for row, col := row-1, col-1; col >= 0 && row >= 0; row, col = row-1, col-1 {
		if arr[row][col] == "Q" {
			return false
		}
	}
	//判断右上角
	for row, col := row-1, col+1; col < n && row >= 0; row, col = row-1, col+1 {
		if arr[row][col] == "Q" {
			return false
		}
	}
	return true
}
