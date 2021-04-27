package tryLock

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewLock(t *testing.T) {
	var counter int
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.UnLock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
