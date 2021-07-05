package main

import "fmt"

/*

	给一个有 n 个结点的有向无环图，找到所有从 0 到 n-1 的路径并输出（不要求按顺序）
	二维数组的第 i 个数组中的单元都表示有向图中 i 号结点所能到达的下一些结点（译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a ）
	空就是没有下一个结点了。

	输入：graph = [[1,2],[3],[3],[]]
	输出：[[0,1,3],[0,2,3]]
	解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3

	链接：https://leetcode-cn.com/problems/all-paths-from-source-to-target
*/
//[[3,1],[4,6,7,2,5],[4,6,3],[6,4],[7,6,5],[6],[7],[]]
func main() {
	graph := [][]int{
		{3, 1},
		{4, 6, 7, 2, 5},
		{4, 6, 3},
		{6, 4},
		{7, 6, 5},
		{6},
		{7},
		{},
	}
	fmt.Println(allPathsSourceTarget(graph))
}

var res [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	path := make([]int, 0)
	res = make([][]int, 0)
	traverse(graph, 0, path)
	return res
}
func traverse(graph [][]int, s int, path []int) {
	//fmt.Printf("%p\n", path)
	path = append(path, s)
	//fmt.Println(res, "res", path, s)
	//fmt.Printf("%p\n", path)
	//fmt.Println("=================")
	if s == len(graph)-1 {
		// TODO append并不是每次都会改变path的地址，只有cap不足的时候才会更改，会导致[0 3 4 5] [0 3 4 6 7] [0 3 4 5 6 7]
		// TODO 这种情况，底层在[0,3,4]的时候地址相同，在append[0,3,4,7]的是地址和[0,3,4]相同，先被改成[0 3 4 6],
		// TODO 最终被改成[0，3，4，5]
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for _, item := range graph[s] {
		traverse(graph, item, path)
	}
}
