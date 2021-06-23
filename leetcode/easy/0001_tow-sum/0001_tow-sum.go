package main

import (
	"fmt"
	"sort"
)

/*
	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
	你可以按任意顺序返回答案。
	https://leetcode-cn.com/problems/two-sum/
*/

func main1() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum2(nums, target))
}

func twoSum(nums []int, target int) []int {
	for index, item := range nums {
		for x_index, x := range nums[index+1:] {
			if target == item+x {
				// go 语言删除数组中间的元素
				//nums = append(nums[0:index+x_index], nums[index+x_index+1:]...)
				return []int{index, index + x_index + 1}
			}
		}
	}
	return []int{}
}
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for index, item := range nums {
		m[item] = index
	}
	fmt.Println(m)
	for i := 0; i < len(nums); i++ {
		b := target - nums[i]
		if num, ok := m[b]; ok && num != i {
			return []int{i, num}
		}
	}
	return []int{}

}

//双指针这里是返回排序后的下标有问题
func twoSum3(nums []int, target int) []int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			left++
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			return []int{left, right}
		}
	}
	return []int{}
}

/* ================*/
type twoSumFind struct {
	arr []int
	dp  map[int]int
}

func NewTwoSumFind() *twoSumFind {
	return &twoSumFind{
		arr: make([]int, 0),
		dp:  make(map[int]int),
	}
}

func (two *twoSumFind) add(num []int) {
	for i := 0; i < len(num); i++ {
		two.dp[num[i]]++
	}
}
func (two *twoSumFind) find(target int) bool {
	for key, value := range two.dp {
		temp := target - key
		if temp == target && value > 1 {
			return true
		}
		_, ok := two.dp[temp]
		if temp != target && ok {
			return true
		}
	}
	return false
}
func main() {
	nums := []int{3, 3, 2, 5}
	target := 6
	dd := NewTwoSumFind()
	dd.add(nums)
	fmt.Println(dd.dp)
	dd.find(target)
}
