package main

/*
	波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
	0 <= n <= 30

*/
import "fmt"

func main() {
	fmt.Println(fib(10))
}
func fib(n int) int {
	var mk = map[int]int{
		1: 1,
		2: 1,
	}

	for i := 3; i <= n; i++ {
		mk[i] = mk[i-1] + mk[i-2]
	}
	return mk[n]
}
