package main

/*
	给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
	不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。



	说明:

	为什么返回数值是整数，但输出的答案是数组呢?

	请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

	你可以想象内部操作如下:

	// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
	int len = removeDuplicates(nums);

	// 在函数里修改输入数组对于调用者是可见的。
	// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
	for (int i = 0; i < len; i++) {
	    print(nums[i]);
	}


	输入：nums = [1,1,2]
	输出：2, nums = [1,2]

	链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
*/
import (
	"fmt"
)

func main() {
	//var  []int{0,0,1,1,1,2,2,3,3,4}
	var arr = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ret := removeDuplicates3(arr)
	fmt.Println(ret)
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
