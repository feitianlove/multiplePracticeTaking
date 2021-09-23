package main

import "fmt"

func main() {
	//d := DifferenceConstruct([]int{1, 2, 3})
	//fmt.Println(d.Diff)
	//d.Increment(0, 2, 1)
	//fmt.Println(d.Result())
}

func corpFlightBookings(bookings [][]int, n int) []int {
	var nums = make([]int, n+1)

	d := DifferenceConstruct(nums)
	for _, item := range bookings {
		d.Increment(item[0], item[1], item[2])
	}
	res := d.Result()
	fmt.Println(res)
	return []int{}
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
