package main

import "fmt"

/*
	根据一棵树的中序遍历与后序遍历构造二叉树。

	注意:
	你可以假设树中没有重复的元素。

	例如，给出

	中序遍历 inorder = [9,3,15,20,7]
	后序遍历 postorder = [9,15,7,20,3]
	返回如下的二叉树：

		3
	   / \
	  9  20
		/  \
	   15   7


	链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder, 0, len(inorder), 0, len(postorder))
}

//	中序遍历 inorder = [9,3,15,20,7]
//	后序遍历 postorder = [9,15,7,20,3]
func build(inorder []int, postorder []int, inStart, inEnd, postStart, postEnd int) *TreeNode {
	fmt.Println(inStart, inEnd, postStart, postEnd)
	//0 5 0 5 3
	//0 1 0 1 9
	if len(inorder[inStart:inEnd]) == 0 {
		return nil
	}
	rootVal := postorder[postEnd-1]
	index := -1
	for i := inStart; i < inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	var root *TreeNode = &TreeNode{
		Val: rootVal,
	}
	leftSize := index - inStart
	root.Left = build(inorder, postorder, inStart, index, postStart, postStart+leftSize)
	root.Right = build(inorder, postorder, index+1, inEnd, postStart+leftSize, postEnd-1)
	return root
}
