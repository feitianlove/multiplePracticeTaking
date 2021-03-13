package main

/*
	罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
	字符          数值
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
	例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
	通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
	I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
	给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

	链接：https://leetcode-cn.com/problems/roman-to-integer

*/
import (
	"fmt"
	"strings"
)

func main() {
	result := romanToInt2("IV")
	fmt.Println(result)
}

func romanToInt(s string) int {
	numSlice := strings.Split(s, "")
	var result []int
	for _, item := range numSlice {
		var temp int
		switch item {
		case "I":
			temp = 1
		case "V":
			temp = 5
		case "X":
			temp = 10
		case "L":
			temp = 50
		case "C":
			temp = 100
		case "D":
			temp = 500
		case "M":
			temp = 1000
		}
		result = append(result, temp)
	}
	var re int
	fmt.Println(result)
	for i := 0; i < len(result)-1; i++ {

		if result[i] < result[i+1] {
			re += -result[i]
		} else {
			re += result[i]
		}

	}
	re += result[len(result)-1]
	return re
}

func romanToInt2(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	current := 0
	prov := 0
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		flag := 1
		current = m[s[i]]
		if current < prov {
			flag = -1
		}
		prov = current
		result += current * flag
	}
	return result
}
