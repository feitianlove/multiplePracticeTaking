package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

func main() {
	//0x76E
	//0xC460
	var a, b string
	for {
		_, err := fmt.Scan(&a, &b)
		if err == io.EOF {
			break
		}
		fmt.Println(translate(a))
		fmt.Println(translate(b))
	}

}
func translate(a string) int {
	totle := 0
	for i := 2; i < len(a); i++ {
		pw := len(a) - i - 1
		switch a[i] {
		case 'a', 'A':
			//temp, _ := strconv.Atoi(string(a[i]))
			totle = totle + int(float64(10)*math.Pow(float64(16), float64(pw)))
		case 'b', 'B':
			totle = totle + int(float64(11)*math.Pow(float64(16), float64(pw)))

		case 'c', 'C':
			totle = totle + int(float64(12)*math.Pow(float64(16), float64(pw)))
		case 'd', 'D':
			totle = totle + int(float64(13)*math.Pow(float64(16), float64(pw)))

		case 'e', 'E':
			totle = totle + int(float64(14)*math.Pow(float64(16), float64(pw)))
		case 'f', 'F':
			totle = totle + int(float64(15)*math.Pow(float64(16), float64(pw)))
		default:
			temp, _ := strconv.Atoi(string(a[i]))
			totle = totle + int(float64(temp)*math.Pow(float64(16), float64(pw)))
		}
	}
	return totle
}
