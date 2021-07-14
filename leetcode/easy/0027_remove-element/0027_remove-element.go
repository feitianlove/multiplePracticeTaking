package main

import "fmt"

func main() {
	var arr = []int{3, 3, 3, 3, 3}
	val := 4
	ret := removeElement2(arr, val)
	fmt.Println(ret)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, length := 0, len(nums)
	for ; i < length; i++ {
		if nums[i] == val {
			temp := nums[i]

			for length > i && nums[length-1] == val {
				length--
			}
			if length == i {
				return i
			}
			nums[i] = nums[length-1]
			nums[length-1] = temp
			length--
		}
	}
	return length
}

//
func removeElement2(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}

	}
	return i
}
