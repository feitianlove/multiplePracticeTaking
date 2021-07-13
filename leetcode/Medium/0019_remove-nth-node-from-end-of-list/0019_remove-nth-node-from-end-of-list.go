package main

/*
	给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
	进阶：你能尝试使用一趟扫描实现吗？

	输入：head = [1,2,3,4,5], n = 2
	输出：[1,2,3,5]
	示例 2：

	输入：head = [1], n = 1
	输出：[]
	示例 3：

	输入：head = [1,2], n = 1
	输出：[1]

	链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
*/
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dp := make([]*ListNode, 0)
	temp := head
	for temp != nil {
		dp = append(dp, temp)
		temp = temp.Next
	}
	delTarget := len(dp) - n
	if delTarget == 0 {
		return dp[delTarget].Next
	}
	dp[delTarget-1].Next = dp[delTarget].Next
	return head
}

//快慢指针

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return slow.Next
	}
	var preSlow *ListNode
	for fast != nil {
		preSlow = slow
		fast = fast.Next
		slow = slow.Next
	}
	preSlow.Next = preSlow.Next.Next
	return head
}
