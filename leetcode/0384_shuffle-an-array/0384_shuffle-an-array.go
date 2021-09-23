package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := Constructor([]int{1, 2, 3})
	dd := s.Shuffle()
	fmt.Println(s.Reset())
	fmt.Println(dd)
}

type Solution struct {
	RawNums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		RawNums: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.RawNums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.RawNums))
	copy(nums, this.RawNums)
	n := len(this.RawNums)
	for i := 0; i < n; i++ {
		randomInt := i + rand.Intn(n-i)

		nums[i], nums[randomInt] = nums[randomInt], nums[i]
	}
	return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
