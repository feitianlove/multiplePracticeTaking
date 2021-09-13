package main

import "fmt"

/*
	编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。


	输入：00000000000000000000000000001011
	输出：3
	解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

	链接：https://leetcode-cn.com/problems/number-of-1-bits

*/
func main() {
	// 1、利用或操作 | 和空格将英文字符转换为小写
	// ' ' 32 0010 0000
	// 'A' 65 0100 0001
	//   	  0110 0001
	// A - a 差了32, | 相当于+32
	fmt.Println(string('a' | ' '))
	fmt.Println(string('A' | ' '))
	//2、利用与操作 & 和下划线将英文字符转换为大写
	// '_' 95 0101 1111
	// 'a' 97 0110 0001
	//        0100 0001
	fmt.Println(string('a' & '_'))
	fmt.Println(string('A' & '_'))
	//3、利用异或操作 ^ 和空格进行英文字符大小写互换
	// 'a' 97 0110 0001
	// ' ' 32 0010 0000
	//		  0100 0001
	fmt.Println(string('a' ^ ' '))
	fmt.Println(string('A' ^ ' '))
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}

func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}
