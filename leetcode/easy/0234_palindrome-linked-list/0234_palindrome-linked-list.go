package main

import "fmt"

/*
	请判断一个链表是否为回文链表。

	示例 1:

	输入: 1->2
	输出: false
	示例 2:

	输入: 1->2->2->1
	输出: true

	https://leetcode-cn.com/problems/palindrome-linked-list/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	t1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	t := &ListNode{
		Val:  2,
		Next: t1,
	}

	var head ListNode = ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: t,
		},
	}
	fmt.Println(isPalindrome2(&head))
}

// 其实就是后续遍历+双指针
var left *ListNode

func isPalindrome(head *ListNode) bool {
	left = head
	return reverse(head)
}

func reverse(head *ListNode) bool {
	if head == nil {
		return true
	}
	//fmt.Println(head.Val)
	res := reverse(head.Next)
	//fmt.Println(head.Val)
	res = res && head.Val == left.Val
	left = left.Next
	return res
}

// 快慢指针优化空间复杂度o(1)
func isPalindrome2(head *ListNode) bool {
	// 1->2-1->nil
	// 1->2->2->1->nil
	if head == nil {
		return true
	}
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		fmt.Println(slow.Val, fast.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	//反转slow开始的
	right := reverse2(slow)
	left := head
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}
	return true
}

func reverse2(head *ListNode) *ListNode {
	//// 1->2-1->nil
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

func (list *ListNode) print() {
	temp := list
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}
