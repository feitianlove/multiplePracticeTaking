package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

}

// 贪心算法
func maxSubArray(nums []int) int {
	result := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > result {
			result = sum
		}
	}
	return result
}

//暴力法
func maxSubArray2(nums []int) int {
	result := math.MinInt32

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > result {
				result = sum
			}
		}
	}
	return result
}

// 动态规划
func maxSubArray3(nums []int) int {
	var dp = make([]int, len(nums))
	var res = math.MinInt32
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
