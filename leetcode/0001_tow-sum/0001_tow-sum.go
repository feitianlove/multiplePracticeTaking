package main

import "fmt"

/*
	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
	你可以按任意顺序返回答案。
	https://leetcode-cn.com/problems/two-sum/
*/
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

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum2(nums, target)
}
