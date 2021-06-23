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
	fmt.Println(threeSum2([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum2([]int{1, 1, 1, 2, 3}))
	fmt.Println(threeSum2([]int{}))
	fmt.Println(threeSum2([]int{0, 0, 0}))

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

//so easy(如果会two sum)
func twoSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	low, height := 0, len(nums)-1
	for low < height {
		left, right := nums[low], nums[height]
		if left+right < target {
			low++
		} else if left+right > target {
			height--
		} else {
			res = append(res, []int{left, right})
			for low < height && left == nums[low] {
				low++
			}
			for low < height && right == nums[height] {
				height--
			}
		}
	}
	return res
}

//有很多重复的, 为了去重复，代码很恶心🤢
func threeSum2(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	m := make(map[[2]int]int)
	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		num := make([]int, len(nums))
		copy(num, nums)
		num = append(num[0:i], num[i+1:]...)
		temp := twoSum(num, target)
		for j := 0; j < len(temp); j++ {
			//fmt.Println(temp)
			//小于说明起那么已经出现过了
			if _, ok := m[[2]int{temp[j][0], temp[j][1]}]; ok {
				continue
			}
			m[[2]int{temp[j][0], temp[j][1]}]++
			if temp[j][1] > nums[i] {
				continue
			}
			tempJ := append(temp[j], nums[i])
			res = append(res, tempJ)
		}
		if i < len(num)-1 && nums[i] == nums[i+1] {
			i++
		}

	}
	return res
}

//a b c
