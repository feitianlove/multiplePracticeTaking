package main

/*

	给定一个链表，判断链表中是否有环。
	如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置
	（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
	如果链表中存在环，则返回 true 。 否则，返回 false 。
	https://leetcode-cn.com/problems/linked-list-cycle/
*/

import (
	"fmt"
	"math"
)

func main() {
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth
	fourth.Next = &second

	fmt.Println(hasCycle(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var dk = make(map[*ListNode]int)
	for head != nil {
		dk[head]++
		if dk[head] != 1 {
			return true
		}
		head = head.Next
	}
	return false
}

//遍历标记
func hasCycle2(head *ListNode) bool {
	for head != nil {
		if head.Val == math.MaxInt32 {
			return true
		}
		head.Val = math.MaxInt32
		head = head.Next
	}
	return false
}
