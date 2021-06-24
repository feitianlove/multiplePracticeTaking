package main

import "fmt"

/*
	给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
	如果数组中不存在目标值 target，返回 [-1, -1]。
	输入：nums = [5,7,7,8,8,10], target = 8
	输出：[3,4]
	链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
*/
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 9
	fmt.Println(searchRange(nums, target))
}
func searchRange(nums []int, target int) []int {
	right := rightSearch(nums, target)
	left := leftSearch(nums, target)
	return []int{left, right}
}

func rightSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var middle int
	for left <= right {
		middle = left + (right-left)/2
		//找右边
		if nums[middle] == target {
			left = middle + 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return left - 1
}

//5, 7, 7, 8, 8, 10
func leftSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var middle int
	for left <= right {
		//println("---------------")
		middle = left + (right-left)/2
		//fmt.Println(left, right, middle)
		//fmt.Println(nums[middle], target)
		//找左边先
		if nums[middle] == target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}
