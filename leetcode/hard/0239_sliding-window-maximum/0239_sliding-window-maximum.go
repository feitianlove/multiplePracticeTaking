package main

import (
	"fmt"
	"math"
)

/*
	给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移
	动一位。返回滑动窗口中的最大值。


	输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
	输出：[3,3,5,5,6,7]
	解释：
	滑动窗口的位置                最大值
	---------------               -----
	[1  3  -1] -3  5  3  6  7       3
	 1 [3  -1  -3] 5  3  6  7       3
	 1  3 [-1  -3  5] 3  6  7       5
	 1  3  -1 [-3  5  3] 6  7       5
	 1  3  -1  -3 [5  3  6] 7       6
	 1  3  -1  -3  5 [3  6  7]      7
	示例 2：

	链接：https://leetcode-cn.com/problems/sliding-window-maximum
*/
func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//[]
	//4
	//nums := []int{-7, -8, 7, 5, 7, 1, 6, 0}
	fmt.Println(maxSlidingWindow2(nums, 4))
}

// 暴力（超出时间限制）
func maxSlidingWindow(nums []int, k int) []int {
	var res = make([]int, 0)
	for i := 0; i < len(nums)-k+1; i++ {
		max := math.MinInt32
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res
}

//使用单调队列
type MonotonicQueue struct {
	queue []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		queue: make([]int, 0),
	}
}

func (queue *MonotonicQueue) Push(n int) {
	for len(queue.queue) != 0 && n > queue.getLast() {
		queue.PopLast()
	}
	queue.queue = append(queue.queue, n)
}

// 获取队尾部
func (queue *MonotonicQueue) getLast() int {
	l := len(queue.queue)
	if l != 0 {
		return queue.queue[l-1]
	} else {
		return -1
	}
}

// 队列尾部出货
func (queue *MonotonicQueue) PopLast() {
	if len(queue.queue) != 0 {
		queue.queue = queue.queue[:len(queue.queue)-1]
	}
}

// 队列头部
func (queue *MonotonicQueue) PollFirst(n int) {
	if queue.getFirst() == n {
		queue.queue = queue.queue[1:]
	}
}

// 获取max,队头，应该是最大
func (queue *MonotonicQueue) getFirst() int {
	if len(queue.queue) != 0 {
		return queue.queue[0]
	} else {
		return -1
	}
}

//使用单调队列
//nums = [-7, -8, 7, 5, 7, 1, 6, 0]
func maxSlidingWindow2(nums []int, k int) []int {
	var res = make([]int, 0)
	var monotonicQueue = NewMonotonicQueue()
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			fmt.Println(monotonicQueue.queue, i)
			monotonicQueue.Push(nums[i])
			fmt.Println(monotonicQueue.queue, i)

		} else {
			monotonicQueue.Push(nums[i])
			fmt.Println(monotonicQueue.queue, i)
			res = append(res, monotonicQueue.getFirst())
			monotonicQueue.PollFirst(nums[i-k+1])
		}
	}
	return res
}
