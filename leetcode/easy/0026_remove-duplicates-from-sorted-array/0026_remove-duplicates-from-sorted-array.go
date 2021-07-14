package main

import (
	"fmt"
)

func main() {
	//var  []int{0,0,1,1,1,2,2,3,3,4}
	var arr = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ret := removeDuplicates3(arr)
	fmt.Println(ret)
}

//TODO removeDuplicates 问题大大的
func removeDuplicates(nums []int) int {

	ret := make(map[int]int)
	for _, item := range nums {
		if _, ok := ret[item]; ok {
			ret[item]++
		} else {
			ret[item] = 1
		}
	}
	fmt.Println(ret)
	//for k, _ := range ret {
	//	fmt.Println(k)
	//}
	return len(ret)
}
func removeDuplicates2(nums []int) int {
	i, j, length := 0, 1, len(nums)
	for ; j < length; j++ {
		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1
}
func removeDuplicates3(nums []int) int {
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[count] = nums[i+1]
			count++
		}
	}
	return count
}
