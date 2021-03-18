package main

import "fmt"

/*
	"【连续最长数字串】
	输入一个字符串，返回其最长的数字子串，若有多个最长的数字子串，则返回最后一个。
	示例1：
	输入：'abcd12345ed125ss123058789'
	输出：'123058789'
	示例2：
	输入：'1643xyyackk254396dd165344'
	输出：'165344'"
*/
func main() {
	fmt.Println(numString("abcd12345ed125ss123058789"))
}
func numString(s string) string {
	temp := ""
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i]-'0' >= 0 && s[i]-'0' <= 9 {
			fmt.Println(string(s[i]))
			temp += string(s[i])
		} else {
			fmt.Println(temp)
			if len(temp) > len(res) {
				res = temp
			}
			temp = ""
		}
	}
	if len(temp) > len(res) {
		return temp
	}
	return res
}
