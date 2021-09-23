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

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 差分数组，对同一个数组的某个区间做加减（频繁）
type Difference struct {
	Diff []int
}

func DifferenceConstruct(num []int) Difference {
	var diff = make([]int, len(num))
	diff[0] = num[0]
	for i := 1; i < len(num); i++ {
		diff[i] = num[i] - num[i-1]
	}
	return Difference{Diff: diff}
}

func (d *Difference) Increment(i, j int, value int) {
	d.Diff[i] += value
	if j+1 < len(d.Diff) {
		d.Diff[j+1] -= value
	}
}

func (d *Difference) Result() []int {
	var res = make([]int, len(d.Diff))
	res[0] = d.Diff[0]
	for i := 1; i < len(d.Diff); i++ {
		res[i] = d.Diff[i] + res[i-1]
	}
	return res
}

// Quick Sort
func QuickSort(num []int, l, r int) []int {
	var res = make([]int, 0)
	if len(num) == 0 {
		return []int{}
	}
	left, right := make([]int, 0), make([]int, 0)
	p := num[l]
	for i := l + 1; i < r; i++ {
		if num[i] < p {
			left = append(left, num[i])
		} else {
			right = append(right, num[i])
		}
	}
	res = append(QuickSort(left, 0, len(left)), p)
	res = append(res, QuickSort(right, 0, len(right))...)
	return res
}
