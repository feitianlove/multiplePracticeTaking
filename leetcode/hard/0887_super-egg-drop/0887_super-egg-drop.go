package main

import (
	"fmt"
)

/*

	给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。
	已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。
	每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，
	则可以在之后的操作中 重复使用 这枚鸡蛋。
	请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？


	输入：k = 1, n = 2
	输出：2
	解释：
	鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。
	否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。
	如果它没碎，那么肯定能得出 f = 2 。
	因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。

链接：https://leetcode-cn.com/problems/super-egg-drop

*/
func main() {
	fmt.Println(superEggDrop(4, 5000))
}

var memo [][]int

func superEggDrop(k int, n int) int {
	memo = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, k+1)
	}
	return dp(k, n)
}

func dp(k int, n int) int {
	if n == 0 {
		return 0
	}
	if k == 1 {
		return n
	}
	if memo[n][k] != 0 {
		return memo[n][k]
	}
	l := 1
	r := n
	for l <= r {
		middle := l + (r-l)/2
		broken := dp(k-1, middle-1)
		unbroken := dp(k, n-middle)
		if broken >= unbroken {
			r = middle - 1
		} else {
			l = middle + 1
		}
	}
	memo[n][k] = 1 + max(dp(k-1, l-1), dp(k, n-l))
	return memo[n][k]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
