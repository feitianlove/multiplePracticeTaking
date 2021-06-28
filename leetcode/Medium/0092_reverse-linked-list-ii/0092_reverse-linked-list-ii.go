package main

import (
	"fmt"
)

/*

	给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

	链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
*/

func main() {
	var head ListNode = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	//dd := reverseList(&head)
	//dd.print()
	//前N个
	//dd := reverseListN(&head, 3)
	//dd.print()
	//Between
	dd := reverseBetween(&head, 2, 3)
	dd.print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//单链表反转
//1->2->3->4->null
//head.next.next= head
//head.next=nil
func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//前N个链表反转
// 需要记录一下head或者end 都可以
var he1 *ListNode

//1->2->3->null
func reverseListN(head *ListNode, n int) *ListNode {
	if n == 1 {
		he1 = head.Next
		return head
	}
	end := reverseListN(head.Next, n-1)
	head.Next.Next = head
	head.Next = he1
	return end
}

//反转n-m之间的链表
//1->2->3->null
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseListN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func (list *ListNode) print() {
	temp := list
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}
