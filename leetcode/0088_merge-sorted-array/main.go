package main

import (
	"fmt"
	"sort"
)

/*
	输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	输出：[1,2,2,3,5,6]
	nums1.length == m + n
	nums2.length == n
	0 <= m, n <= 200
	1 <= m + n <= 200
	-109 <= nums1[i], nums2[i] <= 10

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/merge-sorted-array
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{4, 5, 6}
	n := 3
	merge2(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)

}

//
func merge2(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[n+m-1] = nums1[m-1]
			m--
		} else {
			nums1[n+m-1] = nums2[n-1]
			n--
		}
	}
	if n > 0 {
		for n > 0 {
			nums1[n-1] = nums2[n-1]
			n--
		}
	}
}
