package main

import (
	"fmt"
	"strconv"
)

/*
	你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：
	例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
	列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
	字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。

	输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
	输出：6
	解释：
	可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
	注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
	因为当拨动到 "0102" 时这个锁就会被锁定。

	链接：https://leetcode-cn.com/problems/open-the-lock

*/

func main() {
	//var deadends = []string{"0201", "0101", "0102", "1212", "2002"}
	var deadends = []string{"0201", "0101", "0102", "1212", "2002"}

	fmt.Println(openLock(deadends, "0202"))
	//fmt.Println(openLock2(deadends, "0202"))

}

//自己的
func openLock(deadends []string, target string) int {
	//记录回头路
	var has = make(map[string]int)
	queue := make([]string, 0)
	queue = append(queue, "0000")
	has["0000"]++
	if target == "0000" {
		return 0
	}
	for i := 0; i < len(deadends); i++ {
		has[deadends[i]]++
	}
	tt := 0
	for len(queue) != 0 {
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			temp := queue[i]
			for j := 0; j < 4; j++ {
				e, _ := strconv.Atoi(string(temp[j]))
				//判断是否回头路
				e1 := string(temp[:j]) + plusOne(e) + string(temp[j+1:])
				e2 := string(temp[:j]) + minOne(e) + string(temp[j+1:])
				if e1 == target || e2 == target {
					tt++
					return tt
				}

				if _, ok := has[e1]; !ok {
					if _, ok := has[e1]; !ok {
						has[e1]++
						queue = append(queue, e1)
					}

				}
				if _, ok := has[e2]; !ok {
					if _, ok := has[e2]; !ok {
						has[e2]++
						queue = append(queue, e2)
					}
				}
			}
		}
		queue = queue[qSize:]
		tt++
	}
	return -1
}

func plusOne(n int) string {
	if n == 9 {
		n = 0
	} else {
		n = n + 1
	}
	return strconv.Itoa(n)
}
func minOne(n int) string {
	if n == 0 {
		n = 9
	} else {
		n = n - 1
	}
	return strconv.Itoa(n)
}

//看到大佬的优化
func openLock2(deadends []string, target string) int {
	//记录回头路
	var has = make(map[string]int)
	queue := make([]string, 0)
	queue = append(queue, "0000")
	has["0000"]++
	dk := []int{1, -1}
	if target == "0000" {
		return 0
	}
	for i := 0; i < len(deadends); i++ {
		if deadends[i] == "0000" || deadends[i] == target {
			return -1
		}
		has[deadends[i]]++
	}
	tt := 0
	for len(queue) != 0 {
		tt++
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			temp := queue[i]
			for j := 0; j < 4; j++ {
				for k := 0; k < len(dk); k++ {
					char := string((int(temp[j]-'0')+10+dk[k])%10 + '0')
					e1 := string(temp[:j]) + char + string(temp[j+1:])
					if _, ok := has[e1]; ok {
						continue
					}
					if e1 == target {
						return tt
					}
					has[e1] = 1
					queue = append(queue, e1)
				}
			}
		}
		queue = queue[qSize:]
	}
	return -1
}
