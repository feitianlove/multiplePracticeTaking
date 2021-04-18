package main

import (
	"fmt"
	"sort"
)

/*
	给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
	注意：答案中不可以包含重复的三元组。



	示例 1：

	输入：nums = [-1,0,1,2,-1,-4]
	输出：[[-1,-1,2],[-1,0,1]]
	示例 2：

	输入：nums = []
	输出：[]
	示例 3：j

	输入：nums = [0]
	输出：[]

	链接：https://leetcode-cn.com/problems/3sum
*/

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	p := make(map[int]int)
	m := make(map[[2]int]int)
	fmt.Println(nums)
	for index, item := range nums {
		p[item] = index
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := nums[i] + nums[j]
			if sum > 0 {
				break
			}
			if value, ok := p[-sum]; ok && value > j {
				if _, ok2 := m[[2]int{nums[i], nums[j]}]; !ok2 {
					res = append(res, []int{nums[i], nums[j], 0 - nums[i] - nums[j]})
					m[[2]int{nums[i], nums[j]}] = 1
				}
			}
		}
	}
	return res
}

//a b c
