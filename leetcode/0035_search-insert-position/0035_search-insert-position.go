package main

import "fmt"

func main() {
	var arr = []int{1, 3, 5, 6}
	ret := searchInsert2(arr, 5)
	fmt.Println(ret)
}

// 二分查找
func searchInsert(nums []int, target int) int {
	low, height := 0, len(nums)-1
	for low <= height {
		middle := (low + height) / 2
		switch {
		case target > nums[middle]:
			low = middle + 1
		case target < nums[middle]:
			height = middle - 1
		default:
			return middle
		}
	}
	return height + 1
	//return low
}

//循环搜索
func searchInsert2(nums []int, target int) int {
	i := 0
	for i < len(nums) && nums[i] < target {
		if nums[i] == target {
			return i
		}
		i++
	}
	return i
}
