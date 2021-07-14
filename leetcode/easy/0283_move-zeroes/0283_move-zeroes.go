package main

import "fmt"

/*
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

	输入: [0,1,0,3,12]
	输出: [1,3,12,0,0]
	说明:

	必须在原数组上操作，不能拷贝额外的数组。
	尽量减少操作次数。


	链接：https://leetcode-cn.com/problems/move-zeroes
*/
// TODO 链表数组去重复， 快慢指针
func main() {
	var arr = []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)
}

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	res := 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		} else {
			res++
		}
	}
	for res > 0 {
		nums[len(nums)-res] = 0
		res--
	}
}
