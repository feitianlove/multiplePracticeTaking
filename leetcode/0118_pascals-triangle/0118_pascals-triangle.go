package main

import "fmt"

/*
	给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

*/
func main() {
	fmt.Println(generate2(4))
}

func generate(numRows int) [][]int {
	var result = [][]int{{1}, {1, 1}}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	row := 2
	for numRows != row {
		row++
		var temp []int
		temp = append(temp, 1)
		cur := result[len(result)-1]
		for i := 0; i < len(cur)-1; i++ {
			temp = append(temp, cur[i]+cur[i+1])
		}
		temp = append(temp, 1)
		result = append(result, temp)
	}
	return result
}

func generate2(numRows int) [][]int {
	res := make([][]int, 0)
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		res = append(res, genNext(res[i-1]))
	}
	return res
}

func genNext(p []int) []int {
	res := make([]int, 1, len(p)+1)
	res = append(res, p...)

	for i := 0; i < len(res)-1; i++ {
		res[i] = res[i] + res[i+1]
	}
	return res
}
