package main

import (
	"fmt"
	"sort"
)

/*
	给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
	只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。
	在完成所有删除操作后，请你返回列表中剩余区间的数目。
	输入：intervals = [[1,4],[3,6],[2,8]]
	输出：2
	解释：区间 [3,6] 被区间 [2,8] 覆盖，所以它被删除了
	链接：https://leetcode-cn.com/problems/remove-covered-intervals
*/
func main() {
	intervals := [][]int{
		//{1, 4}, {3, 6}, {2, 8},
		{1, 4}, {3, 6}, {2, 5},
	}
	fmt.Println(removeCoveredIntervals(intervals))
}

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	//	[[1 4] [2 8] [3 6]]
	fmt.Println(intervals)
	left, right := intervals[0][0], intervals[0][1]
	var res int
	for i := 1; i < len(intervals); i++ {
		inte := intervals[i]
		//情况一 完全包含
		if left <= inte[0] && right >= inte[1] {
			res++
		}
		//|____|
		//  |____|
		// 情况二 找到相交区间，合并
		if right >= inte[0] && right <= inte[1] {
			right = inte[1]
		}
		//|____|
		//  			|____|
		// 情况三 没有交集
		if right < inte[0] {
			left = inte[0]
			right = inte[1]
		}
		fmt.Println(left, right)

	}
	return len(intervals) - res
}
