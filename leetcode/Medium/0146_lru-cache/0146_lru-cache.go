package main

import "fmt"

/*

	运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
	实现 LRUCache 类：

	LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
	int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
	void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该
	在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。


	进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？

	示例：

	输入
	["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	输出
	[null, null, null, 1, null, -1, null, -1, 3, 4]
	解释
	LRUCache lRUCache = new LRUCache(2);
	lRUCache.put(1, 1); // 缓存是 {1=1}
	lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
	lRUCache.get(1);    // 返回 1
	lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.get(2);    // 返回 -1 (未找到)
	lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.get(1);    // 返回 -1 (未找到)
	lRUCache.get(3);    // 返回 3
	lRUCache.get(4);    // 返回 4

	链接：https://leetcode-cn.com/problems/lru-cache
*/
//[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
func main() {
	lru := Constructor(2)
	lru.Put(1, 0)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))

	lru.Put(3, 3)
	fmt.Println(lru.HashMap)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.HashMap)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}

type LRUCache struct {
	HashMap map[int]*Link
	Link    *Link
	size    int //容量
}

type Link struct {
	Val  int
	Key  int
	Prev *Link
	Next *Link
}

// 需要加一个头尾的指针指向链表
var head, tail *Link

func Constructor(capacity int) LRUCache {
	head = &Link{}
	tail = &Link{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		HashMap: make(map[int]*Link),
		Link:    nil,
		size:    capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.HashMap[key]; !ok {
		return -1
	} else {
		// 将其删除之后加入到尾部
		this.deleteLink(node)
		this.AddLast(node)
		return node.Val
	}

}

func (this *LRUCache) Put(key int, value int) {
	var node = &Link{
		Val: value,
		Key: key,
	}
	// 如果已经存在了更新
	if existNode, ok := this.HashMap[key]; ok {
		this.deleteLink(existNode)
		this.AddLast(node)
		this.HashMap[key] = node
		return
	}
	if this.Size() < this.size {
		this.HashMap[key] = node
		this.AddLast(node)
	} else {
		// 将该数据删除后加到链表尾部
		this.deleteFirstLink()
		this.AddLast(node)
		this.HashMap[key] = node
	}
}

// 在链表尾部添加一个元素
func (this *LRUCache) AddLast(x *Link) {
	x.Next = tail
	x.Prev = tail.Prev
	tail.Prev.Next = x
	tail.Prev = x
}

//删除头部的元素
func (this *LRUCache) deleteFirstLink() {
	if head.Next == tail {
		return
	}
	first := head.Next
	this.deleteLink(first)
	delete(this.HashMap, first.Key)
}

//删除链表中的 x 节点（x 一定存在）
func (this *LRUCache) deleteLink(x *Link) {
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
}

//返回当前的容量
func (this *LRUCache) Size() int {
	return len(this.HashMap)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
