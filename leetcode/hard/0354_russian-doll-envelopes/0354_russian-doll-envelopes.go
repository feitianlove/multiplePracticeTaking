package main

import (
	"fmt"
	"sort"
)

/*
	给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
	当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
	请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

	输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
	输出：3
	解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。

	输入：envelopes = [[1,1],[1,1],[1,1]]
	输出：1


	链接：https://leetcode-cn.com/problems/russian-doll-envelopes
*/
func main() {
	var envelopes = [][]int{
		{5, 4}, {6, 4}, {6, 7}, {2, 3},
	}
	//var envelopes = [][]int{
	//	{1, 1}, {1, 1}, {1, 1},
	//}
	fmt.Println(maxEnvelopes(envelopes))
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})
	fmt.Println(envelopes)
	var dp = make([]int, len(envelopes))
	var res int
	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] {
				if envelopes[i][1] > envelopes[j][1] {
					dp[i] = max(dp[j]+1, dp[i])
				}
			}
		}
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

// TODO
//这个解法的关键在于，对于宽度 w 相同的数对，要对其高度 h 进行降序排序。因为两个宽度相同的信封不能相互包含的，
//逆序排序保证在 w 相同的数对中最多只选取一个。
//先对宽度 w 进行升序排序，如果遇到 w 相同的情况，则按照高度 h 降序排序。之后把所有的 h 作为一个数组，
//在这个数组上计算 LIS 的长度就是答案转换成最长递增子序列的问题了😊
