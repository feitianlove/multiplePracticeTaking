package main

import "fmt"

/*

	给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。

	如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。



	示例 1：

	输入：n = 1
	输出：true
	解释：20 = 1

	链接：https://leetcode-cn.com/problems/power-of-two
*/

func main() {
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(12))

}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}
