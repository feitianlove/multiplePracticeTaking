package main

import "fmt"

/*
	给定一个已按照 升序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
	函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，所以答案数组应当满足
	1 <= answer[0] < answer[1] <= numbers.length 。
	你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

	输入：numbers = [2,7,11,15], target = 9
	输出：[1,2]
	解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
	示例 2：

	输入：numbers = [2,3,4], target = 6
	输出：[1,3]
	示例 3：

	输入：numbers = [-1,0], target = -1
	输出：[1,2]

	链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
*/
func main() {
	numbers := []int{2, 3, 4}
	fmt.Println(twoSum2(numbers, 6))
}

func twoSum(numbers []int, target int) []int {
	dp := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		dp[numbers[i]] = i
	}
	for i := 0; i < len(numbers); i++ {
		if value, ok := dp[target-numbers[i]]; ok && value != i {
			return []int{i + 1, value + 1}
		}
	}
	return []int{}
}

//使用二分查找

func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{-1, -1}
}
