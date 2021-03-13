package main

import (
	"fmt"
)

/*
	设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
	push(x) —— 将元素 x 推入栈中。
	pop() —— 删除栈顶的元素。
	top() —— 获取栈顶元素。
	getMin() —— 检索栈中的最小元素。

	链接：https://leetcode-cn.com/problems/min-stack
*/

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.stack)
	fmt.Println(minStack.GetMin())
	fmt.Println(minStack.stack)
	minStack.Pop()
	fmt.Println(minStack.stack)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())

}

/** initialize your data structure here. */

type item struct {
	min, x int
}
type MinStack struct {
	stack []item
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{
		min: min,
		x:   x,
	})
}

func (this *MinStack) Pop() {

	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {

	data := this.stack[len(this.stack)-1].x
	this.stack = this.stack[:len(this.stack)-1]
	return data
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// TODO 也可以用双栈来实现
type SMinStack struct {
	data []int
	min  []int
}
