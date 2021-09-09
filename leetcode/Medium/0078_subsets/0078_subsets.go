package main

import "fmt"

/*
	给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
	解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

	输入：nums = [1,2,3]
	输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

	链接：https://leetcode-cn.com/problems/subsets
*/
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	var track = make([]int, 0)
	backtrack(nums, 0, track)
	return result
}

func backtrack(nums []int, start int, track []int) {
	// 用copy
	var temp = make([]int, len(track))
	copy(temp, track)

	result = append(result, temp)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1, track)
		track = track[:len(track)-1]
	}

}
