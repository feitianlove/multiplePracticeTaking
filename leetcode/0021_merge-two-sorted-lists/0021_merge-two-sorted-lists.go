package main

import "fmt"

func main() {
	//输入：l1 = [1,2,4], l2 = [1,3,4]
	//输出：[1,1,2,3,4,4]

	var first, firstone, firsttwo ListNode
	first.Val = 1
	firstone.Val = 2
	firsttwo.Val = 4
	first.Next = &firstone
	firstone.Next = &firsttwo

	var second, secondone, secondtwo ListNode
	second.Val = 1
	secondone.Val = 3
	secondtwo.Val = 4
	second.Next = &secondone
	secondone.Next = &secondtwo
	node := mergeTwoLists2(&first, &second)
	for {
		fmt.Print(node.Val, "->")
		node = node.Next
		if node.Next == nil {
			fmt.Println(node.Val)
			break
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head

}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	}
}
