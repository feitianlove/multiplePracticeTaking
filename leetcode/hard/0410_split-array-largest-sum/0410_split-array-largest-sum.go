package main

import "fmt"

/*
	给定一个非负整数数组 nums 和一个整数 m ，你需要将这个数组分成 m 个非空的连续子数组。
	设计一个算法使得这 m 个子数组各自和的最大值最小。

	输入：nums = [7,2,5,10,8], m = 2
	输出：18
	解释：
	一共有四种方法将 nums 分割为 2 个子数组。 其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
	因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。

	链接：https://leetcode-cn.com/problems/split-array-largest-sum
*/
func main() {
	arr := []int{7, 2, 5, 10, 8}
	fmt.Println(splitArray(arr, 2))
}

func splitArray(nums []int, m int) int {
	left := getMax(nums)
	right := getSum(nums)
	for left <= right {
		middle := left + (right-left)/2
		if split(nums, middle) > m {
			left = middle + 1
		} else if split(nums, middle) < m {
			right = middle - 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func split(nums []int, max int) int {
	sum := 0
	res := 1
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > max {
			res++
			sum = nums[i]
		} else {
			sum = sum + nums[i]
		}
	}
	return res
}
func getSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
func getMax(nums []int) int {
	max := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
