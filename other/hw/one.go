package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a = strings.Split(input.Text(), " ")
	}
	//arr := []string{"i", "am", "student"}
	stringR(a)
	fmt.Println(a)
}

func two() {
}

func stringR(arr []string) {
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[len(arr)-i-1]
		arr[len(arr)-i-1] = arr[i]
		arr[i] = temp
	}
}
