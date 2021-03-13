package main

import (
	"fmt"
	"strconv"
)

/*
	输入: a = "1010", b = "1011"
	输出: "10101"
	每个字符串仅由字符 '0' 或 '1' 组成。
	1 <= a.length, b.length <= 10^4
	字符串如果不是 "0" ，就都不含前导零。
*/
func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary2(a, b))
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	length := len(a)
	aarr := transToInt(a, length)
	barr := transToInt(b, length)
	return makeString(add(aarr, barr))
}

func transToInt(s string, length int) []int {
	result := make([]int, length)
	ls := len(s)
	for i, b := range s {
		result[length-ls+i] = int(b - '0')

	}
	return result
}

func add(a, b []int) []int {
	result := make([]int, len(a))
	length := len(a)
	for i := length - 1; i > 0; i-- {
		//fmt.Println(a[i], b[i], (a[i]+b[i])%2)
		result[i] = (a[i] + b[i]) % 2
		if a[i]+b[i] >= 2 {
			a[i-1]++
		} else {
		}
	}
	if a[0]+b[0] >= 2 {
		temp := (a[0] + b[0]) / 2
		result[0] = (a[0] + b[0]) % 2
		result = append([]int{temp}, result...)
	} else {
		result[0] = a[0] + b[0]
	}
	return result
}

func makeString(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}
	return string(bytes)
}

// 直接
func addBinary2(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	result := ""
	flag := 0
	for i >= 0 || j >= 0 {
		intA, intB := 0, 0
		if i >= 0 {
			intA = int(a[i] - '0')
		}
		if j >= 0 {
			intB = int(b[j] - '0')
		}
		current := intA + intB + flag
		flag = 0
		if current >= 2 {
			flag = current / 2
			current = current % 2
		}
		cur := strconv.Itoa(current)
		result = cur + result
		i--
		j--
	}
	if flag == 1 {
		result = "1" + result
	}
	return result
}
