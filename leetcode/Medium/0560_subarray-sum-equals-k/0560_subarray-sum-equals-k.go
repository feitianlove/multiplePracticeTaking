package main

import "fmt"

/*
	给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。

	输入：nums = [1,1,1], k = 2
	输出：2

	输入：nums = [1,2,3], k = 3
	输出：2

	链接：https://leetcode-cn.com/problems/subarray-sum-equals-k

*/
func main() {
	//fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	//fmt.Println(subarraySum([]int{1}, 0))
	//fmt.Println(subarraySum([]int{-1, -1, 1}, 0))
	//fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))

}

// 不能使用双指针 ， 值回小于0 , 正确的是前缀数组
func subarraySum(nums []int, k int) int {
	//sum := make([]int, len(nums)+1)
	//for i := 0; i < len(nums); i++ {
	//	sum[i+1] = sum[i] + nums[i]
	//}
	var res int
	var mp = make(map[int]int)
	mp[0] = 1
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		des := sum - k
		if value, ok := mp[des]; ok {
			res += value
		}
		mp[sum]++
	}
	return res
}
