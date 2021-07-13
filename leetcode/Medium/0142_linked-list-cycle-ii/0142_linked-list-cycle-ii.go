package main

/*
	给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
	为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，
	pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

	链接：https://leetcode-cn.com/problems/linked-list-cycle-ii

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}

	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
