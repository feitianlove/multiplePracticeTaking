package main

import "fmt"

/*

	在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有
	一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋
	将自动报警。

	计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
	输入: [3,2,3,null,3,null,1]

		 3
		/ \
	   2   3
		\   \
		 3   1

	输出: 7
	解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.

	链接：https://leetcode-cn.com/problems/house-robber-iii
*/

func main() {

	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(rob(&root))
	fmt.Println(rob2(&root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	dd := make(map[*TreeNode]int)
	return dp(root, dd)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//同house_robber_i 和house_robber_II的思路的思路，（--丝滑的发现超出时间限制了🤢--）
func dp(root *TreeNode, dd map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if value, ok := dd[root]; ok {
		return value
	}
	//抢
	doIt := 0
	if root.Left == nil {
		doIt = root.Val
	} else {
		doIt = root.Val + dp(root.Left.Left, dd) + dp(root.Left.Right, dd)
	}
	if root.Right != nil {
		doIt += dp(root.Right.Left, dd) + dp(root.Right.Right, dd)
	}
	//不抢
	notDo := dp(root.Left, dd) + dp(root.Right, dd)
	dd[root] = max(doIt, notDo)
	return dd[root]
}

// 更舒服的代码
func rob2(root *TreeNode) int {
	res := dp2(root)
	return max(res[0], res[1])
}

func dp2(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dp2(root.Left)
	right := dp2(root.Right)
	//抢
	rob := root.Val + left[0] + right[0]
	//不抢
	notRob := max(left[0], left[1]) + max(right[1], right[0])
	return []int{notRob, rob}
}
