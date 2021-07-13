package main

import (
	"fmt"
	"math/rand"
)

/*
	给定一个包含 [0，n) 中不重复整数的黑名单 blacklist ，写一个函数从 [0, n) 中返回一个不在 blacklist 中的随机整数。
	对它进行优化使其尽量少调用系统方法 Math.random() 。

	提示:

	1 <= n <= 1000000000
	0 <= blacklist.length < min(100000, N)
	[0, n) 不包含 n ，详细参见 interval notation 。
	示例 1：

	输入：
	["Solution","pick","pick","pick"]
	[[1,[]],[],[],[]]
	输出：[null,0,0,0]

	链接：https://leetcode-cn.com/problems/random-pick-with-blacklist
*/
func main() {
	nums := []int{0, 3}
	s := Constructor(4, nums)
	fmt.Println(s.ValOfIndex)
	fmt.Println(s.Sz)
	fmt.Println(s.Pick())
	//fmt.Println(s.Pick())
	//fmt.Println(s.Pick())
	//fmt.Println(s.Pick())
	//fmt.Println(s.Pick())
	//fmt.Println(s.Pick())

}

type Solution struct {
	Sz         int
	ValOfIndex map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	s := Solution{
		Sz:         n - len(blacklist),
		ValOfIndex: make(map[int]int),
	}
	max := n - len(blacklist)
	for i := 0; i < len(blacklist); i++ {
		s.ValOfIndex[blacklist[i]] = 1
	}
	last := n - 1
	for i := 0; i < len(blacklist); i++ {
		if blacklist[i] >= max {
			continue
		}
		// 1,2,3 [0,1]
		for {
			_, ok := s.ValOfIndex[last]
			if ok {
				last--
			} else {
				break
			}
		}
		s.ValOfIndex[blacklist[i]] = last
		last--
	}
	return s
}

func (this *Solution) Pick() int {
	index := rand.Intn(this.Sz)
	fmt.Println(index, "index")
	if value, ok := this.ValOfIndex[index]; ok {
		return value
	}
	return index
}
