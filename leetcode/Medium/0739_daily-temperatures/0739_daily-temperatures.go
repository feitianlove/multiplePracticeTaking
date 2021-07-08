package main

import "fmt"

/*
	请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用
	 0 来代替。例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

	提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

	链接：https://leetcode-cn.com/problems/daily-temperatures

*/
func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// TODO 用dp了 ，使用append([]int{stack[len(stack)-1] - i}, res...) 超出时间限制了
	dp := make(map[int]int)
	var stack = make([]int, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		fmt.Println(stack)
		for !(len(stack) == 0) && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			dp[i] = 0
		} else {
			dp[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	//[1 1 4 2 1 1 0 0]
	for key, value := range dp {
		res[key] = value
	}
	return res
}
