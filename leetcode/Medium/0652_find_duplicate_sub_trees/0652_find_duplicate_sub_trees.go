package main

import (
	"fmt"
	"strconv"
)

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
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(findDuplicateSubtrees(&root))
}

//二叉树的序列化
var res []*TreeNode

func traverse(root *TreeNode, hash map[string]int) string {
	if root == nil {
		return "#"
	}
	left := traverse(root.Left, hash)
	right := traverse(root.Right, hash)
	freq := left + "," + right + "," + strconv.Itoa(root.Val)
	if value, ok := hash[freq]; ok && value == 1 {
		res = append(res, root)
	}
	hash[freq]++
	return freq
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	hash := make(map[string]int)
	res = make([]*TreeNode, 0)
	traverse(root, hash)
	return res
}
