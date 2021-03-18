package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}))
	fmt.Println(compress([]byte{'a', 'a', 'a', 'a', 'a', 'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}))

}
func compress(chars []byte) int {
	temp := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			temp++
		} else {
			if temp == 1 {
			} else {
				li := chars[i:]
				for in, num := range strconv.Itoa(temp) {
					chars = append(chars[:i-temp+1+in], byte(num))
				}
				chars = append(chars, li...)
				i = i - temp + 2
			}
			temp = 1
		}
	}
	if temp != 1 {
		chars = chars[:len(chars)-temp+1]
		for _, num := range strconv.Itoa(temp) {
			chars = append(chars, byte(num))
		}
	}
	fmt.Println(chars)
	return len(chars)
}
