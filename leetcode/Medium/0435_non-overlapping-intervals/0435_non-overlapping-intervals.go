package main

import (
	"fmt"
	"sort"
)

/*
	给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
	注意:

	可以认为区间的终点总是大于它的起点。
	区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
	示例 1:

	输入: [ [1,2], [2,3], [3,4], [1,3] ]
	输出: 1
	解释: 移除 [1,3] 后，剩下的区间没有重叠。

	示例 2:
	输入: [ [1,2], [1,2], [1,2] ]
	输出: 2
	解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。

	示例 3:
	输入: [ [1,2], [2,3] ]
	输出: 0

	链接：https://leetcode-cn.com/problems/non-overlapping-intervals
*/
func main() {
	var intervals = [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{1, 3},
		//{1, 2},
		//{1, 2},
		//{1, 2},
	}
	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	var res int
	for i := 1; i < len(intervals); i++ {
		it := intervals[i]
		// 包含
		// ｜________｜
		//    ｜__｜
		if left <= it[0] && right >= it[1] {
			left = it[0]
			right = it[1]
			res++
		}
		// 相交
		// ｜________｜
		//    ｜_________｜
		if right > it[0] && right < it[1] {
			res++
		}
		if right <= it[0] {
			left = it[0]
			right = it[1]
		}
		fmt.Println(left, right, res)
	}
	return res
}

//
