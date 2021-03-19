package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO 第一种
	//var a, b string
	//for {
	//	_, err := fmt.Scan(&a, &b)
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(a, b)
	//}
	//TODO 第二种
	var a, b string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a = strings.Split(input.Text(), " ")[0]
		b = strings.Split(input.Text(), " ")[1]
		fmt.Println(a + b)
	}
}
