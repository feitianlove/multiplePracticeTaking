package main

import (
	"fmt"
)

/*
	给定一个 没有重复 数字的序列，返回其所有可能的全排列。

	示例:

	输入: [1,2,3]
	输出:
	[
	  [1,2,3],
	  [1,3,2],
	  [2,1,3],
	  [2,3,1],
	  [3,1,2],
	  [3,2,1]
	]
	链接：https://leetcode-cn.com/problems/permutations

*/
func main() {
	fmt.Println(permute([]int{1, 2}))
}

//递归
func permute1(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		tempArr := make([]int, len(nums)-1)
		copy(tempArr[:i], nums[:i])
		copy(tempArr[i:], nums[i+1:])
		arr := permute1(tempArr)
		for _, v := range arr {
			res = append(res, append(v, nums[i]))
		}
	}
	return res
}

//回溯算法
var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	arr := make([]int, 0)
	visited := make(map[int]bool)
	dfs(nums, 0, arr, visited)
	return res
}

func dfs(nums []int, index int, arr []int, visited map[int]bool) {
	if index == len(nums) {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == false {
			visited[i] = true
			arr = append(arr, nums[i])
			dfs(nums, index+1, arr, visited)
			visited[i] = false
			arr = arr[:len(arr)-1]
		}
	}
}
