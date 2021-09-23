package main

import (
	"fmt"
	"time"
)

func main() {
	num := []int{3, 3, 3, 3, 3, 3, 3, 3, 3}
	fmt.Println(findKthLargest(num, 1))
}

func findKthLargest(nums []int, k int) int {
	lo, hi := 0, len(nums)-1
	index := len(nums) - k
	for lo <= hi {
		p := partition(nums, lo, hi)
		fmt.Println(p)
		time.Sleep(time.Second)
		if p > index {
			hi = p - 1
		} else if p < index {
			lo = p + 1
		} else {
			return nums[p]
		}
	}
	return -1
}

func QuickSort(num []int, lo, hi int) []int {
	if lo >= hi {
		return num
	}
	// 通过交换元素构建分界点索引 p
	p := partition(num, lo, hi)
	QuickSort(num, lo, p-1)
	QuickSort(num, p+1, hi)
	return num
}
func partition(num []int, lo, hi int) int {
	if lo == hi {
		return lo
	}
	pivot := num[lo]
	i, j := lo+1, hi
	for {

		for num[i] <= pivot {
			if i == hi {
				break
			}
			i++
		}
		for num[j] > pivot {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		num[i], num[j] = num[j], num[i]
	}
	num[j], num[lo] = num[lo], num[j]
	return j
}
