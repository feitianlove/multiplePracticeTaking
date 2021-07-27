package main

import (
	"fmt"
	"sort"
)

/*
	你将会获得一系列视频片段，这些片段来自于一项持续时长为 T 秒的体育赛事。这些片段可能有所重叠，也可能长度不一。
	视频片段 clips[i] 都用区间进行表示：开始于 clips[i][0] 并于 clips[i][1] 结束。我们甚至可以对这些片段自由地再剪辑，例如片段 [0, 7] 可
	以剪切成 [0, 1] + [1, 3] + [3, 7] 三部分。我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, T]）。返回所需
	片段的最小数目，如果无法完成该任务，则返回 -1 。


	输入：clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
	输出：3
	解释：
	我们选中 [0,2], [8,10], [1,9] 这三个片段。
	然后，按下面的方案重制比赛片段：
	将 [1,9] 再剪辑为 [1,2] + [2,8] + [8,9] 。
	现在我们手上有 [0,2] + [2,8] + [8,10]，而这些涵盖了整场比赛 [0, 10]。
	示例 2：

	输入：clips = [[0,1],[1,2]], T = 5
	输出：-1
	解释：
	我们无法只用 [0,1] 和 [1,2] 覆盖 [0,5] 的整个过程。


	链接：https://leetcode-cn.com/problems/video-stitching
*/

func main() {
	clips := [][]int{
		{0, 2},
		{4, 6},
		{8, 10},
		{1, 9},
		{1, 5},
		{5, 9},
	}
	fmt.Println(videoStitching(clips, 10))
}

func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})
	left := clips[0][0]
	if left != 0 {
		return -1
	}
	var count int = 0
	var i int = 0
	var res int
	var currentEnd = 0
	for i < len(clips) && clips[i][0] <= currentEnd {
		for i < len(clips) && clips[i][0] <= currentEnd {
			res = max(res, clips[i][1])
			i++
		}
		currentEnd = res
		count++
		if currentEnd >= time {
			return count
		}
	}
	if currentEnd < time {
		return -1
	}
	return count

}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
