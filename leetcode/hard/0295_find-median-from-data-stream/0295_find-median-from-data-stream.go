package main

import "sort"

/*
	中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

	例如，

	[2,3,4] 的中位数是 3

	[2,3] 的中位数是 (2 + 3) / 2 = 2.5

	设计一个支持以下两种操作的数据结构：

	void addNum(int num) - 从数据流中添加一个整数到数据结构中。
	double findMedian() - 返回目前所有元素的中位数。
	示例：

	addNum(1)
	addNum(2)
	findMedian() -> 1.5
	addNum(3)
	findMedian() -> 2

	链接：https://leetcode-cn.com/problems/find-median-from-data-stream

*/
func main() {

}

//TODO 需要优化 leetcode超出时间限制
// 可以使用插入排序直接解，但是寻找插入位置还要但是插入操作需要搬移数据最差O(n)

type MedianFinder struct {
	large []int
	small []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		large: make([]int, 0),
		small: make([]int, 0),
	}
}

// 要保证large和small的差不能大于1，并且small的值必须小于large的最大值
func (this *MedianFinder) AddNum(num int) {
	if len(this.small) >= len(this.large) {
		this.small = append(this.small, num)
		sort.Ints(this.small)
		temp := this.small[len(this.small)-1]
		this.small = this.small[:len(this.small)-1]
		this.large = append(this.large, temp)
	} else {
		this.large = append(this.large, num)
		sort.Ints(this.large)
		temp := this.large[len(this.large)-1]
		this.large = this.large[1:len(this.small)]
		this.small = append(this.small, temp)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.small)
	sort.Ints(this.large)
	if len(this.small) == len(this.large) {
		return (float64(this.small[len(this.small)-1]) + float64(this.large[0])/2)
	} else if len(this.large) > len(this.small) {
		return float64(this.large[0])
	} else {
		return float64(this.small[len(this.small)-1])
	}
}
