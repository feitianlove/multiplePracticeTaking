package main

import (
	"fmt"
	"math"
)

/*
	f(x) 是 x! 末尾是 0 的数量。（回想一下 x! = 1 * 2 * 3 * ... * x，且 0! = 1 ）
	例如， f(3) = 0 ，因为 3! = 6 的末尾没有 0 ；而 f(11) = 2 ，因为 11!= 39916800 末端有 2 个 0 。给定 K，找出多少个非负整数 x ，
	能满足 f(x) = K 。

	输入：K = 0
	输出：5
	解释：0!, 1!, 2!, 3!, and 4! 均符合 K = 0 的条件。

	链接：https://leetcode-cn.com/problems/preimage-size-of-factorial-zeroes-function

*/

func main() {
	fmt.Println(math.Pow(10, 9) < float64(trailingZeroes(math.MaxInt32*10)))
	//fmt.Println(preimageSizeFZF(5))
}

// 二分查找[0, math.MaxUint32] 右区间，值为k
func preimageSizeFZF(k int) int {
	res := rightBound(int64(k)) - leftBound(int64(k)) + 1
	return int(res)
}

func rightBound(k int64) int64 {
	var left, right int64 = 0, math.MaxInt32 * 10
	for left < right {
		middle := left + (right-left)/2
		if trailingZeroes(middle) > k {
			right = middle
		} else if trailingZeroes(middle) < k {
			left = middle + 1
		} else {
			left = middle + 1
		}
	}
	return left - 1
}
func leftBound(k int64) int64 {
	var left, right int64 = 0, math.MaxInt32 * 10
	for left < right {
		middle := left + (right-left)/2
		if trailingZeroes(middle) > k {
			right = middle
		} else if trailingZeroes(middle) < k {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return right
}

func trailingZeroes(n int64) int64 {
	var res int64
	var div int64 = 5
	for n >= div {
		res += n / div
		div *= 5
	}
	return res
}
