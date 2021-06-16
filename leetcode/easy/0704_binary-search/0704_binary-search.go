package main

import "fmt"

/*
	给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

	输入: nums = [-1,0,3,5,9,12], target = 9
	输出: 4
	解释: 9 出现在 nums 中并且下标为 4

	链接：https://leetcode-cn.com/problems/binary-search
*/

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search2(nums, target))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if target < nums[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

//{-1, 0, 3, 5, 9, 12}
func search2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if target < nums[middle] {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return -1
}
