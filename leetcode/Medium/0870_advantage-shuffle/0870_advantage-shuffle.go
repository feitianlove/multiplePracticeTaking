package main

import (
	"fmt"
	"sort"
)

/*

	给定两个大小相等的数组 A 和 B，A 相对于 B 的优势可以用满足 A[i] > B[i] 的索引 i 的数目来描述。
	返回 A 的任意排列，使其相对于 B 的优势最大化。

	输入：A = [2,7,11,15], B = [1,10,4,11]
	输出：[2,11,7,15]
	示例 2：

	输入：A = [12,24,8,32], B = [13,25,32,11]
	输出：[24,32,8,12]

	链接：https://leetcode-cn.com/problems/advantage-shuffle

*/

func main() {
	num1 := []int{12, 24, 8, 32}
	num2 := []int{13, 25, 32, 11}
	fmt.Println(advantageCount(num1, num2))
}

func advantageCount(nums1 []int, nums2 []int) []int {
	var res = make([]int, len(nums1))
	sort.Ints(nums1)
	var dp = make(map[int]int)
	arr := make([][2]int, 0)
	for i := 0; i < len(nums2); i++ {
		arr = append(arr, [2]int{i, nums2[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	temp := make([]int, len(nums2))
	sort.Ints(temp)
	left, right := 0, len(nums1)-1
	for i := len(temp) - 1; i >= 0; i-- {
		idx := dp[temp[i]]
		if temp[i] < nums1[right] {
			res[idx] = nums1[right]
			right--
		} else {
			res[idx] = nums1[left]
			left++
		}
	}
	return res
}
