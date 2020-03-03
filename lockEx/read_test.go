package lockex

import (
	"sync"
	"testing"
	"time"
)

var m *sync.RWMutex

func Test_read(t *testing.T) {
	m = new(sync.RWMutex)

	// 多个同时读
	go read(1)
	go read(2)

	time.Sleep(2 * time.Second)
}

func read(i int) {
	println(i, "read start")

	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read over")
}
