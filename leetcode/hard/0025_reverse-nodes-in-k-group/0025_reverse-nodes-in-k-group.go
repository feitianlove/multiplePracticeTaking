package main

/*
	给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
	k 是一个正整数，它的值小于或等于链表的长度。
	如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

	链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
*/

import "fmt"

func main() {
	t := &ListNode{
		Val:  3,
		Next: nil,
	}
	var head ListNode = ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: t,
		},
	}

	// 单链表反转
	//dd := reverse(&head)
	//dd.print()
	//前n个
	//dd := reverseN(&head, 2)
	//dd.print()
	//到b之前的反转
	//dd := reverseN2(&head, t)
	//dd.print()
	//N3
	//dd := reverseN3(&head, t)
	//dd.print()
	dd := reverseKGroup(&head, 2)
	dd.print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 不用递归的单链表反转
//1 -> 2 ->3 -> nill
func reverse(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur, next := head, head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//1 -> 2 ->3 -> nill 2
func reverseN(head *ListNode, n int) *ListNode {
	var pre *ListNode = nil
	cur, next := head, head
	var h *ListNode = head
	for n >= 1 {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		n--
	}
	h.Next = next
	return pre
}

func reverseN2(head *ListNode, b *ListNode) *ListNode {
	var pre *ListNode = nil
	cur, next := head, head
	var h *ListNode = head
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	h.Next = next
	return pre
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverseN3(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// reverseKGroup使用
func reverseN3(head *ListNode, b *ListNode) *ListNode {
	var pre *ListNode = nil
	cur, next := head, head
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func (list *ListNode) print() {
	temp := list
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}
