package main

/*
	给你二叉树的根结点 root ，请你将它展开为一个单链表：
	展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
	展开后的单链表应该与二叉树 先序遍历 顺序相同。

	链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	flatten(&root)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left
	var p *TreeNode = root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

}
