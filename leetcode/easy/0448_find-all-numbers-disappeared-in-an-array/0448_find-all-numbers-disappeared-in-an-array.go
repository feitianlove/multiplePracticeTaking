package main

import (
	"fmt"
	"math"
)

/*
	给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式
	返回结果。

	输入：nums = [4,3,2,7,8,2,3,1]
	输出：[5,6]

	链接：https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array

*/
func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}), "res")
}

// 是人就会的
func findDisappearedNumbers1(nums []int) []int {
	var dp = make(map[int]int)
	var res = make([]int, 0)
	for i := 0; i < len(nums); i++ {
		dp[nums[i]]++
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := dp[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}

// 4, 3, 2, 7, 8, 2, 3, 1
// 把对应索引的值变成负数，下一次是负数就是重复的 痛645
func findDisappearedNumbers(nums []int) []int {
	var res = make([]int, 0)
	for i := 0; i < len(nums); i++ {
		index := int(math.Abs(float64(nums[i])))
		if nums[index-1] > 0 {
			nums[index-1] *= -1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
