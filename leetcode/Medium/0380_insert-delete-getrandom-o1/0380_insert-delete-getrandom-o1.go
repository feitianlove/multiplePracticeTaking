package main

import (
	"fmt"
	"math/rand"
)

/*
	设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。
	insert(val)：当元素 val 不存在时，向集合中插入该项。
	remove(val)：元素 val 存在时，从集合中移除该项。
	getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
	示例 :

	// 初始化一个空的集合。
	RandomizedSet randomSet = new RandomizedSet();

	// 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	randomSet.insert(1);

	// 返回 false ，表示集合中不存在 2 。
	randomSet.remove(2);

	// 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
	randomSet.insert(2);

	// getRandom 应随机返回 1 或 2 。
	randomSet.getRandom();

	// 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
	randomSet.remove(1);

	// 2 已在集合中，所以返回 false 。
	randomSet.insert(2);

	// 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
	randomSet.getRandom();

	链接：https://leetcode-cn.com/problems/insert-delete-getrandom-o1

*/
func main() {
	randomizedSet := Constructor()
	fmt.Println(randomizedSet.Insert(1))
	fmt.Println(randomizedSet.Insert(2))
	fmt.Println(randomizedSet.Insert(3))
	fmt.Println(randomizedSet.ValToIndex)
	fmt.Println(randomizedSet.List)
	fmt.Println(randomizedSet.Remove(2))
	fmt.Println(randomizedSet.Remove(2))

	fmt.Println(randomizedSet.ValToIndex)
	fmt.Println(randomizedSet.List)

}

type RandomizedSet struct {
	List       []int
	ValToIndex map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		List:       make([]int, 0),
		ValToIndex: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	index := len(this.List)
	if _, ok := this.ValToIndex[val]; ok {
		return false
	}
	this.List = append(this.List, val)
	this.ValToIndex[val] = index
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	l := len(this.List)
	if index, ok := this.ValToIndex[val]; ok {
		temp := this.List[l-1]
		//this.List[l-1] = this.List[index]
		this.List[index] = temp
		this.ValToIndex[temp] = index
		delete(this.ValToIndex, val)
		this.List = this.List[:l-1]
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.List[rand.Intn(len(this.List))]
}
