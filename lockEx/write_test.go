package lockex

import (
	"sync"
	"testing"
	"time"
)

// var m *sync.Mutex

func Test_write(t *testing.T) {
	m = new(sync.RWMutex)

	// 同時寫
	go write(1)
	go write(2)

	time.Sleep(2 * time.Second)
}

func write(i int) {
	println(i, "write start")

	m.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()

	println(i, "write over")
}
