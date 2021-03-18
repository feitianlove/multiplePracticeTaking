package main

import (
	"fmt"
	"sort"
)

/*
	珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉
	，然后这一小时内不会再吃更多的香蕉

	https://leetcode-cn.com/problems/koko-eating-bananas/
*/
func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	max := piles[len(piles)-1]
	left, right := 1, max
	for left < right {
		middle := left + (left+right)/2
		fmt.Println(left, right, middle)
		if judge(piles, middle, h) == true {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return left
}

func judge(piles []int, speed int, h int) bool {
	totle := 0
	for i := 0; i < len(piles); i++ {
		totle = totle + piles[i]/speed
		if piles[i]%speed > 0 {
			totle = totle + 1
		}
	}
	return totle > h
}
