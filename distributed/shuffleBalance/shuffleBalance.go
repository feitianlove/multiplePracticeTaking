package shuffleBalance

import (
	"fmt"
	"math/rand"
	"time"
)

func Shuffle(slice []interface{}) {
	for i := len(slice); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)

		slice[lastIdx], slice[idx] = slice[idx], slice[lastIdx]
	}
	//b := rand.Perm(n)
}

func Request(params map[string]interface{}) error {
	// 读取endpoint
	var indexes = []interface{}{
		"100.69.62.1:3232",
		"100.69.62.32:3232",
		"100.69.62.42:3232",
		"100.69.62.81:3232",
		"100.69.62.11:3232",
		"100.69.62.113:3232",
		"100.69.62.101:3232",
	}
	Shuffle(indexes)
	maxRetryTimes := 3
	var err error
	for i := 0; i < maxRetryTimes; i++ {
		err = apiRequest(params, indexes[i].(string))
		fmt.Println(params, indexes[i])
		if err == nil {
			break
		}
		time.Sleep(10 * time.Microsecond)
	}
	if err != nil {
		return err
	}
	return nil
}

//实际的业务请求
func apiRequest(params map[string]interface{}, endpoint string) error {
	return nil
}
