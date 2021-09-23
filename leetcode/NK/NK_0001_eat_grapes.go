package main

/*
	https://www.nowcoder.com/questionTerminal/14c0359fb77a48319f0122ec175c9ada
*/
import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eatGrapes(1, 2, 6))
}

func eatGrapes(a, b, c int) int {
	arr := []int{a, b, c}
	sum := a + b + c
	sort.Ints(arr)

	if arr[0]+arr[1] > arr[2] {
		// 向上取证  sum +（n-1）/n
		return sum + 2/3
	}
	if 2*(arr[0]+arr[1]) < arr[2] {
		return (arr[3] + 1) / 2
	}
	return (sum + 2) / 3
}
