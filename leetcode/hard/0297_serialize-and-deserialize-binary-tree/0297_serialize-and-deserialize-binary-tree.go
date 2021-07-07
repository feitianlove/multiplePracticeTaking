package main

/*
	序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，
	采取相反方式重构得到原数据。请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可
	以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

	输入：root = [1,2,3,null,null,4,5]
	输出：[1,2,3,null,null,4,5]


	链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree

*/
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	cc := Constructor()
	fmt.Println(cc.serialize(root))
	fmt.Println(cc.deserialize(cc.serialize(root)))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
	res []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.res = strings.Split(data, ",")
	fmt.Println(this.res)
	return this.dfsDeserialize()
}

func (this *Codec) dfsDeserialize() *TreeNode {
	node := this.res[0]
	this.res = this.res[1:]
	if node == "#" {
		return nil
	}
	value, _ := strconv.Atoi(node)
	return &TreeNode{
		Val:   value,
		Left:  this.dfsDeserialize(),
		Right: this.dfsDeserialize(),
	}
}
