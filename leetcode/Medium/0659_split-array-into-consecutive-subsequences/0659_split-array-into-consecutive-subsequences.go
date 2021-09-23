package main

import "fmt"

/*
	给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个长度至少为 3 的子序列，其中每个子序列都由连续整数组成。
	如果可以完成上述分割，则返回 true ；否则，返回 false 。

	输入: [1,2,3,3,4,5]
	1, 2, 3
	3, 4, 5

	链接：https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences
*/

func main() {
	fmt.Println(isPossible([]int{1, 2, 3}))
}

func isPossible(nums []int) bool {
	var freq = make(map[int]int)
	var need = make(map[int][][]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if freq[num] == 0 {
			continue
		} else if freq[num] > 0 && len(need[num]) > 0 {
			freq[num]--
			temp := make([]int, len(need[num][0]))
			copy(temp, need[num][0])
			need[num] = need[num][1:]
			temp = append(temp, num)
			need[num+1] = append(need[num+1], temp)
		} else if freq[num] > 0 && freq[num+1] > 0 && freq[num+2] > 0 {
			freq[num]--
			freq[num+1]--
			freq[num+2]--
			need[num+3] = append(need[num+3], []int{num, num + 1, num + 2})
		} else {
			return false
		}
	}
	return true
}
