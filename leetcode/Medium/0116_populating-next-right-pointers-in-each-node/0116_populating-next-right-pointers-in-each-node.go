package main

import "fmt"

/*

	给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

	struct Node {
	  int val;
	  Node *left;
	  Node *right;
	  Node *next;
	}
	填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
	初始状态下，所有 next 指针都被设置为 NULL。



链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node

*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	root := Node{
		Val: 0,
		Left: &Node{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &Node{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(root)
	fmt.Println(connect(&root))
}

func connect(root *Node) *Node {

	if root == nil {
		return nil
	}
	connectTowNode(root.Left, root.Right)
	return root
}
func connectTowNode(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTowNode(node1.Left, node1.Right)
	connectTowNode(node2.Left, node2.Right)
	// 相邻节点的链接
	connectTowNode(node1.Right, node2.Left)
}
