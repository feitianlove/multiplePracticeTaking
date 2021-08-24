package main

import "fmt"

/*
	有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

	现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

	求所能获得硬币的最大数量。



	示例 1：
	输入：nums = [3,1,5,8]
	输出：167
	解释：
	nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
	coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/burst-balloons
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}

func maxCoins(nums []int) int {
	//重构一下数组
	var NewNums = []int{1}
	nums = append(nums, 1)
	NewNums = append(NewNums, nums...)
	n := len(NewNums)
	// 定义dp  i,j 从i->j 之间 获得的最高分(i,j)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 状态方程 k 表示最后戳破的气球
	// res = max(res, dp[i][k] + dp[k][j] + NewNums[i]*NewNums[k]*NewNums[j])

	for i := n; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+NewNums[i]*NewNums[k]*NewNums[j])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
