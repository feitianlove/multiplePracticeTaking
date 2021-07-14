package main

import "fmt"

/*
	输入: 1->1->2
	输出: 1->2
*/
func main() {
	head := ListNode{
		Val: 1,
	}
	one := ListNode{
		Val: 1,
	}
	two := ListNode{
		Val: 2,
	}
	head.Next = &one
	one.Next = &two
	//deleteDuplicates(&head).print()
	deleteDuplicates(&head)
	//head.print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

//快慢指针
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head.Next, head
	for fast != nil {

		if fast.Val == slow.Val {
			fast = fast.Next
			continue
		}

		slow.Next = fast
		slow = slow.Next
		fast = fast.Next
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := head
	for temp.Next != nil {

		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return head
}

//双指针
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	q := head.Next
	for p.Next != nil {
		if p.Val == q.Val {
			if q.Next == nil {
				p.Next = nil
			} else {
				p.Next = q.Next
				q = q.Next
			}
		} else {
			p = p.Next
			q = q.Next
		}
	}
	return head
}
