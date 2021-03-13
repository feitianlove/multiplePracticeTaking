package main

import (
	"fmt"
	"math"
)

/*
	输入: 8
	输出: 2
	说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
*/
func main() {
	fmt.Println(mySqrt(8))

}

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	low, height := 0, x

	for low+1 < height {
		middle := (low + height) / 2
		switch {
		case middle*middle > x:
			height = middle
		case middle*middle < x:
			low = middle
		default:
			return middle
		}
	}
	return low
}
func mySqrt2(x int) int {
	return int(math.Sqrt(float64(x)))
	//result := math.Floor(math.Sqrt(float64(x)))
}

// 牛顿迭代
func mySqrt3(x int) int {
	result := x
	for result*result > x {
		result = (result + x/result) / 2
	}
	return result
}

// 二分法另外一种
func mySqrt4(x int) int {
	left := 1
	right := x
	for left <= right {
		mid := (left + right) / 2
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left*left <= x {
		return left
	} else {
		return left - 1
	}
}
