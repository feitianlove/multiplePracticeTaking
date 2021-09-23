package main

import "fmt"

/*
	给你一个整数数组 arr ，请使用 煎饼翻转 完成对数组的排序。

	一次煎饼翻转的执行过程如下：

	选择一个整数 k ，1 <= k <= arr.length
	反转子数组 arr[0...k-1]（下标从 0 开始）
	例如，arr = [3,2,1,4] ，选择 k = 3 进行一次煎饼翻转，反转子数组 [3,2,1] ，得到 arr = [1,2,3,4] 。

	以数组形式返回能使 arr 有序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在 10 * arr.length 范围内的有效答案都将被判断为正确。

	输入：[3,2,4,1]
	输出：[4,2,4,3]
	解释：
	我们执行 4 次煎饼翻转，k 值分别为 4，2，4，和 3。
	初始状态 arr = [3, 2, 4, 1]
	第一次翻转后（k = 4）：arr = [1, 4, 2, 3]
	第二次翻转后（k = 2）：arr = [4, 1, 2, 3]
	第三次翻转后（k = 4）：arr = [3, 2, 1, 4]
	第四次翻转后（k = 3）：arr = [1, 2, 3, 4]，此时已完成排序。


*/

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 1, 4}))
}

var res []int

func pancakeSort(arr []int) []int {
	res = make([]int, 0)
	sort(arr, len(arr))
	return res
}

func sort(arr []int, n int) {
	if n == 1 {
		return
	}
	var max, maxIndex int

	for i := 0; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}
	if maxIndex == n-1 {
		sort(arr, n-1)
		return
	}
	//第一次反转将最大的转到数组0
	reverse(arr, 0, maxIndex)
	res = append(res, maxIndex+1)
	//第二次反转将最大的转到数组尾部
	reverse(arr, 0, n-1)
	res = append(res, n)
	sort(arr, n-1)

}

func reverse(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
