package lockex

import (
	"fmt"
	"sync"
	"testing"
)

type Service struct {
	sync.Mutex
}

func Test_mutex(t *testing.T) {
	svc := Service{}
	count := 0
	times := 1000
	wg := new(sync.WaitGroup)
	wg.Add(times)
	call := func() {
		defer wg.Done()
		defer svc.Unlock()
		svc.Lock()
		count++
	}
	for i := 0; i < times; i++ {
		go call()
	}
	wg.Wait()
	fmt.Println(count)
}
