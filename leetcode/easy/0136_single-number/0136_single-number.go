package main

import (
	"fmt"
	"sort"
)

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

输入: [2,2,1]
输出: 1
链接：https://leetcode-cn.com/problems/single-number

*/
func main() {
	var arr = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(arr))
}

/*hash方法*/
func singleNumber3(nums []int) int {
	md := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := md[nums[i]]; ok {
			md[nums[i]] = md[nums[i]] + 1
		} else {
			md[nums[i]] = 1
		}
	}
	for k, v := range md {
		if v == 1 {
			return k
		}
	}
	return -1
}

/*排序遍历*/

func singleNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		if i+1 == len(nums) {
			return nums[i]
		}
		if nums[i+1] != nums[i] {
			return nums[i]
		}
	}
	return -1
}

//TODO 异或 https://github.com/feitianlove/go-leetcode/blob/master/source/0101-0200/0136_single-number/1-%E5%BC%82%E6%88%96.go
//异或
func singleNumber(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
