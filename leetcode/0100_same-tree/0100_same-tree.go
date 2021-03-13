package main

import "fmt"

/*
 	给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
	如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
	输入：p = [1,2,3], q = [1,2,3]
	输出：true
	https://leetcode-cn.com/problems/same-tree/
*/
func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond

	second := TreeNode{Val: 1}
	secondfirst := TreeNode{Val: 2}
	secondSecond := TreeNode{Val: 3}
	second.Left = &secondfirst
	second.Right = &secondSecond
	fmt.Println(isSameTree(&root, &second))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	var queueP, queueQ []*TreeNode
	queueP = append(queueP, p)
	queueQ = append(queueQ, q)
	for len(queueQ) > 0 && len(queueQ) > 0 {
		tempQ := queueQ[0]
		queueQ = queueQ[1:]
		tempP := queueP[0]
		queueP = queueP[1:]
		if tempP.Val != tempQ.Val {
			return false
		}
		if (tempP.Left == nil && tempQ.Left != nil) || (tempP.Left != nil && tempQ.Left == nil) {
			return false
		}
		if tempQ.Left != nil {
			queueP = append(queueP, tempP.Left)
			queueQ = append(queueQ, tempQ.Left)
		}
		if (tempP.Right == nil && tempQ.Right != nil) || tempP.Right != nil && tempQ.Right == nil {
			return false
		}
		if tempQ.Right != nil {
			queueP = append(queueP, tempP.Right)
			queueQ = append(queueQ, tempQ.Right)
		}
	}
	return true
}
