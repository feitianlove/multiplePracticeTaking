package main

import (
	"fmt"
)

// 使用栈结构实现
type stack []rune

func (s *stack) push(data rune) {
	*s = append(*s, data)
}
func (s *stack) pop() (rune, bool) {
	if len(*s) > 0 {
		res := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return res, true
	}
	return 0, false
}
func main() {
	s := "()"
	ret := isValid3(s)
	fmt.Println(ret)
}

var match = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	st := new(stack)
	for _, item := range s {
		switch item {
		case '(', '[', '{':
			st.push(item)
		case ')', ']', '}':
			ret, ok := st.pop()
			if !ok || ret != match[item] {
				return false
			}
		}
	}
	if len(*st) > 0 {
		return false
	}
	return true
}

// 使用数组实现
func isValid2(s string) bool {
	if s == "" {
		return false
	}
	var length int
	stack := make([]rune, len(s))
	var m = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, item := range s {
		switch item {
		case '(', '{', '[':
			stack[length] = item
			length++
		case ')', ']', '}':
			if length <= 0 {
				return false
			}
			if m[item] != stack[length-1] {
				fmt.Println(m[item], stack[length-1])
				return false
			} else {
				length--
			}

		}
	}
	fmt.Println(length)
	if length != 0 {
		return false
	}
	return true
}

func isValid3(s string) bool {
	if s == "" {
		return false
	}
	var stack []rune
	var m = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, item := range s {
		switch item {
		case '(', '{', '[':
			stack = append(stack, item)
		case ')', ']', '}':
			if len(stack) <= 0 {
				return false
			}
			if m[item] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
