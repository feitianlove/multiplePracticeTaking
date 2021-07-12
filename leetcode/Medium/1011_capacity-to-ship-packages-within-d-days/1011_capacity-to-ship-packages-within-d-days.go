package main

import "fmt"

/*
	传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。
	传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
	返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。


	输入：weights = [1,2,3,4,5,6,7,8,9,10], D = 5
	输出：15
	解释：
	船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
	第 1 天：1, 2, 3, 4, 5
	第 2 天：6, 7
	第 3 天：8
	第 4 天：9
	第 5 天：10

	请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。

	链接：https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days
*/

func main() {
	weights := []int{1, 2, 3, 1, 1}
	fmt.Println(shipWithinDays(weights, 4))
}
func shipWithinDays(weights []int, days int) int {
	// 注意： left最小载重就是weights中的最大， right肯定是所有之合了
	left := 0
	right := 0
	for i := 0; i < len(weights); i++ {
		left = max(left, weights[i])
		right += weights[i]
	}
	//[]
	for left <= right {
		middle := left + (right-left)/2
		if f(weights, middle) < days {
			right = middle - 1
		} else if f(weights, middle) > days {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}
func f(weights []int, cap int) int {
	var day = 0
	for i := 0; i < len(weights); {
		var temp = cap
		for i < len(weights) {
			if temp < weights[i] {
				break
			} else {
				temp = temp - weights[i]
			}
			i++
		}
		day++
	}
	return day
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
