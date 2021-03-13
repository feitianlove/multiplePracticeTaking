package main

import (
	"fmt"
)

/*
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
1 <= digits.length <= 100
0 <= digits[i] <= 9
*/
func main() {
	fmt.Println(plusOne2([]int{8, 9, 9, 9}))
}

func plusOne(digits []int) []int {
	carry := 0 // 进位标志

	if digits[len(digits)-1]+1 >= 10 {

		carry = (digits[len(digits)-1] + 1) / 10
		digits[len(digits)-1] = (digits[len(digits)-1] + 1) % 10
	} else {
		digits[len(digits)-1] = (digits[len(digits)-1] + 1) % 10
		return digits
	}

	for i := len(digits) - 2; i >= 0; i-- {
		fmt.Println(digits[i], carry)
		if digits[i]+carry >= 10 {
			carry = (digits[i] + carry) / 10
			digits[i] = (digits[i] + carry) % 10
		} else {
			digits[i] = (digits[i] + carry) % 10
			carry = 0
			break
		}
	}
	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}

func plusOne2(digits []int) []int {
	length := len(digits)
	digits[length-1]++
	for i := length - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] = digits[i] - 10
		digits[i-1]++
	}
	if digits[0] >= 10 {
		digits[0] = digits[0] - 10
		digits = append([]int{1}, digits...)
	}
	return digits
}
