package leetcode

// 1。 去除字符串中飞字符
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

// 2。栈实现
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

// 满二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
