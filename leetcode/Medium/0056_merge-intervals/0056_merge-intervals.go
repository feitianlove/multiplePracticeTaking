package main

import (
	"fmt"
	"sort"
)

/*
	以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，
	该数组需恰好覆盖输入中的所有区间。

	输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
	输出：[[1,6],[8,10],[15,18]]
	解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].


	链接：https://leetcode-cn.com/problems/merge-intervals
*/

func main() {
	intervals := [][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}

	//intervals := [][]int{
	//	{1, 4}, {5, 6},
	//}
	fmt.Println(merge(intervals))
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		inte := intervals[i]
		//第一种情况 不需要处理
		//	｜_____｜
		//｜___________｜
		//
		if left < inte[0] && right > inte[1] {

		}
		//第二种情况
		//	｜_____｜
		//      ｜___________｜
		if right >= inte[0] && right <= inte[1] {
			//如果有合并去掉上一个
			res = res[:len(res)-1]
			right = inte[1]
			res = append(res, []int{left, right})
		}
		if right < inte[0] {
			//第三种情况
			//	｜_____｜
			//     			｜___________｜
			left = inte[0]
			right = inte[1]
			res = append(res, []int{left, right})
		}
	}
	return res
}
