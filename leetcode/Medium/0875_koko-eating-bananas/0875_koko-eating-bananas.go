package main

import (
	"fmt"
)

/*
	珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉
	，然后这一小时内不会再吃更多的香蕉
	输入: piles = [3,6,7,11], H = 8
	输出: 4
	https://leetcode-cn.com/problems/koko-eating-bananas/
*/
func main() {
	piles := []int{312884470}
	fmt.Println(minEatingSpeed(piles, 968709470))
}

//[312884470]
//968709470

func minEatingSpeed(piles []int, h int) int {
	left, right := 0, 1000000000+1
	for left < right {
		middle := left + (right-left)/2
		// 吃太快了
		if middle < 1 {
			return 1
		}
		if eatBananaHours(piles, middle) < h {
			right = middle
		} else if eatBananaHours(piles, middle) > h {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}

func eatBananaHours(piles []int, speed int) int {
	var res int
	for i := 0; i < len(piles); i++ {
		res += piles[i] / speed
		if piles[i]%speed > 0 {
			res++
		}
	}
	return res
}
