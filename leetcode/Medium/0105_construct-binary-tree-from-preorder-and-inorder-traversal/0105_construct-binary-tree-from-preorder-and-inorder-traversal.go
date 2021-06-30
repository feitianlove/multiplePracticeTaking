package main

/*
	根据一棵树的前序遍历与中序遍历构造二叉树。
	注意:
	你可以假设树中没有重复的元素。

	例如，给出

	前序遍历 preorder = [3,9,20,15,7]
	中序遍历 inorder = [9,3,15,20,7]
	返回如下的二叉树：

		3
	   / \
	  9  20
		/  \
	   15   7
	链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, len(preorder), 0, len(inorder))
}

//	前序遍历 preorder = [3,9,20,15,7]
//	[9] [20,15,7]

//	中序遍历 inorder = [9,3,15,20,7]
//  [9] [15,20,7]
func build(preorder []int, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if len(inorder[inStart:inEnd]) == 0 {
		return nil
	}
	rootVal := preorder[preStart]
	index := -1
	for i := inStart; i < inEnd; i++ {
		if rootVal == inorder[i] {
			index = i
			break
		}
	}
	var root *TreeNode = &TreeNode{
		Val: rootVal,
	}
	leftSize := index - inStart
	root.Left = build(preorder, inorder, preStart+1, preStart+leftSize, inStart, index)
	root.Right = build(preorder, inorder, preStart+leftSize+1, preEnd, index+1, inEnd)
	return root
}
