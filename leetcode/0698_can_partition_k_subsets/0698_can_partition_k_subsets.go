package main

import (
	"fmt"
	"sort"
)

/*
	给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。

	输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
	输出： True
	说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。

	1 <= k <= len(nums) <= 16
	0 < nums[i] < 10000

	链接：https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets
*/
func main() {
	//fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1}, 1))
	//fmt.Println(canPartitionKSubsets1([]int{1, 1, 1, 1, 2, 2, 2, 2}, 4))

}

// 典型的回溯算法DFS
// 1、 我们站在nums的角度，每个数字都要选择进入到K个桶中的一个 超时

func canPartitionKSubsets1(nums []int, k int) bool {
	n := len(nums)
	var sum int
	if k > n {
		return false
	}
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	sort.Ints(nums)
	target := sum / k
	bucket := make([]int, k)
	return backtrack1(nums, 0, bucket, target)
}

func backtrack1(nums []int, index int, bucket []int, target int) bool {
	if index == len(nums) {
		for i := 0; i < len(bucket); i++ {
			if bucket[i] != target {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(bucket); i++ {
		// 大于target
		if bucket[i]+nums[index] > target {
			continue
		}
		// 加入bucket
		bucket[i] = bucket[i] + nums[index]
		if backtrack1(nums, index+1, bucket, target) {
			return true
		}
		bucket[i] -= nums[index]
	}
	return false
}

//2、如果我们切换到这 k 个桶的视角，对于每个桶，都要遍历 nums 中的 n 个数字，然后选择是否将当前遍历到的数字装进自己这个桶里

func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	var sum int
	if k > n {
		return false
	}
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	sort.Ints(nums)
	target := sum / k
	var used = make([]bool, len(nums))
	return backtrack(nums, 0, 0, k, target, used)
}

func backtrack(nums []int, index int, bucket, k, target int, used []bool) bool {
	if k == 0 {
		return true
	}
	if bucket == target {
		return backtrack(nums, 0, 0, k-1, target, used)
	}
	for i := index; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if nums[i]+bucket > target {
			continue
		}
		used[i] = true
		bucket = bucket + nums[i]

		if backtrack(nums, i+1, bucket, k, target, used) {
			return true
		}
		used[i] = false
		bucket -= nums[i]
	}
	return false
}
