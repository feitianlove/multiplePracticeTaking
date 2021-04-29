package shuffleBalance

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {

	var cnt1 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		var sl = []interface{}{0, 1, 2, 3, 4, 5, 6}
		Shuffle(sl)
		cnt1[sl[0].(int)]++
	}
	fmt.Println(cnt1)
}
