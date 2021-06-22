package main

import (
	"fmt"
)

/*
	给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi]
	而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。返回这 两个区间列表的交集 。

	形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。
	两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。

	输入：firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
	输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]



	链接：https://leetcode-cn.com/problems/interval-list-intersections

*/

func main() {
	//firstList := [][]int{
	//	{0, 2}, {5, 10}, {13, 23}, {24, 25},
	//}
	//secondList := [][]int{
	//	{1, 5}, {8, 12}, {15, 24}, {25, 26},
	//}
	//
	//firstList := [][]int{
	//	{5, 10},
	//}
	//secondList := [][]int{
	//	{3, 10},
	//}

	firstList := [][]int{
		{4, 6}, {7, 8}, {10, 17},
	}
	secondList := [][]int{
		{5, 10},
	}
	fmt.Println(intervalIntersection(firstList, secondList))
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0)
	if len(firstList) == 0 || len(secondList) == 0 {
		return res
	}
	for i := 0; i < len(firstList); i++ {
		left, right := firstList[i][0], firstList[i][1]
		for j := 0; j < len(secondList); j++ {
			if left == secondList[j][1] {
				res = append(res, []int{left, left})
			}
			//第一种情况
			if left <= secondList[j][0] && right >= secondList[j][1] {
				res = append(res, []int{secondList[j][0], secondList[j][1]})
				continue
			}
			if left >= secondList[j][0] && right <= secondList[j][1] {
				res = append(res, []int{left, right})
				continue
			}
			//第二种情况
			// 		|_________|
			// |_________|
			if right >= secondList[j][0] && right < secondList[j][1] {
				res = append(res, []int{secondList[j][0], right})
			}
			if left < secondList[j][1] && right >= secondList[j][1] {
				res = append(res, []int{left, secondList[j][1]})
			}
			//第三种情况
			if right < secondList[j][0] {
				break
			}
		}

	}
	return res

}
