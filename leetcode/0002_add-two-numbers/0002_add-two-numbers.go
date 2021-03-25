package main

import "fmt"

/*
	给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
	请你将两个数相加，并以相同形式返回一个表示和的链表。
	你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	输入：l1 = [2,4,3], l2 = [5,6,4]
	输出：[7,0,8]
	解释：342 + 465 = 807

	链接：https://leetcode-cn.com/problems/add-two-numbers
*/

func main() {
	first := ListNode{Val: 2}
	first1 := ListNode{Val: 4}
	first2 := ListNode{Val: 3}
	first.Next = &first1
	first1.Next = &first2

	second := ListNode{Val: 5}
	second1 := ListNode{Val: 6}
	second2 := ListNode{Val: 4}
	second.Next = &second1
	second1.Next = &second2

	fmt.Println(addTwoNumbers(&first, &second))
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	res := &ListNode{
		Val: sum % 10,
	}
	if sum >= 10 {
		l1.Next = addTwoNumbers(l1.Next, &ListNode{Val: 1})
	}
	res.Next = addTwoNumbers(l1.Next, l2.Next)
	return res
}
