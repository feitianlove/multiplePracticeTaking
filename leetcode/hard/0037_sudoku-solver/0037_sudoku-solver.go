package main

import "fmt"

/*
	编写一个程序，通过填充空格来解决数独问题。
	数独的解法需 遵循如下规则：

	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
	数独部分空格内已填入了数字，空白格用 '.' 表示。



	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/sudoku-solver
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
}

// 八皇后一个思路 DFS
func solveSudoku(board [][]byte) {
	dd := backtrack(board, 0, 0)
	fmt.Println(dd)
}
func backtrack(board [][]byte, row int, col int) bool {
	if col == 9 {
		return backtrack(board, row+1, 0)
	}
	if row == 9 {
		return true
	}
	if board[row][col] != '.' {
		return backtrack(board, row, col+1)
	}
	for i := '1'; i <= '9'; i++ {
		if !isValid(board, row, col, byte(i)) {
			continue
		}
		board[row][col] = byte(i)
		if backtrack(board, row, col+1) {
			return true
		}
		board[row][col] = '.'
	}
	return false
}

func isValid(board [][]byte, row, col int, ch byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == ch {
			return false
		}
		if board[row][i] == ch {
			return false
		}
		// 注意这里看题： 向下取整
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == ch {
			return false
		}
	}
	return true
}
