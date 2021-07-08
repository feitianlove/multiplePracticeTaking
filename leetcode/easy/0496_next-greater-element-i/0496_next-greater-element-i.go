package main

import "fmt"

/*
	给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
	请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
	nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

	输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
	输出: [-1,3,-1]
	解释:
		对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
		对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
		对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。

	链接：https://leetcode-cn.com/problems/next-greater-element-i

*/

func main() {
	nums2 := []int{1, 2, 3, 4}
	nums1 := []int{2, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

//
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var mp = make(map[int]int)
	var res = make([]int, len(nums1))
	var stack = make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		fmt.Println(stack)

		for !(len(stack) == 0) && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			mp[nums2[i]] = -1
		} else {
			mp[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i := 0; i < len(nums1); i++ {
		res[i] = mp[nums1[i]]
	}
	return res
}
