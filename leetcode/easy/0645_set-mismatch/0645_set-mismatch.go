package main

import (
	"fmt"
	"math"
)

/*
	集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且
	有一个数字重复 。给定一个数组 nums 代表了集合 S 发生错误后的结果。
	请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

	链接：https://leetcode-cn.com/problems/set-mismatch
*/
func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}

func findErrorNums(nums []int) []int {
	var n = len(nums)
	var dump = -1
	for i := 0; i < n; i++ {
		index := int(math.Abs(float64(nums[i])))
		if nums[index-1] < 0 {
			dump = int(math.Abs(float64(nums[i])))
		} else {
			nums[index-1] *= -1
		}
	}
	var miss = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			miss = i + 1
		}
	}
	return []int{dump, miss}
}
