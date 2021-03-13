package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a "
	fmt.Println(lengthOfLastWord2(s))
}

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	strArr := strings.Split(s, " ")
	return len(strArr[len(strArr)-1])
}

//倒序遍历
func lengthOfLastWord2(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	result := 0
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if result > 0 {
				return result
			}
			continue
		}
		result++
	}
	return result
}
