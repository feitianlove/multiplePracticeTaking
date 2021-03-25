package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

// 暴力通过不了。。。。
func maxArea(height []int) int {
	max := 0
	for i := 1; i < len(height); i++ {
		water := 0
		for j := i - 1; j >= 0; j-- {
			water = (i - j) * min(height[i], height[j])
			if water > max {
				max = water
			}
		}
	}
	return max
}

// 双指针
func maxArea2(height []int) int {
	res := 0
	i := 0
	j := len(height) - 1
	for i < j {
		water := (j - i) * min(height[i], height[j])
		if res < water {
			res = water
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}
