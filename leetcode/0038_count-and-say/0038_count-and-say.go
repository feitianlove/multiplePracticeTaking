package main

import (
	"fmt"
	"strconv"
)

func main() {
	ret := countAndSay2(2)
	fmt.Println(ret)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	ret := ""
	i := 0
	temp := 1
	for i < len(str)-1 {
		if str[i] == str[i+1] {
			temp++
		} else {
			ch := fmt.Sprintf("%c", str[i])
			ret += strconv.Itoa(temp) + ch
			temp = 1
		}
		i++
	}
	ch := fmt.Sprintf("%c", str[i])
	ret += strconv.Itoa(temp) + ch

	return ret
}

func countAndSay2(n int) string {
	if n == 1 {
		return "1"
	}
	strs := countAndSay(n - 1)
	result := make([]byte, 0, len(strs)*2)
	i, j := 0, 1
	for i < len(strs) {
		for j < len(strs) && strs[i] == strs[j] {
			j++
		}
		result = append(result, byte(j-i+'0'))
		result = append(result, strs[i])
		i = j
	}
	return string(result)
}
