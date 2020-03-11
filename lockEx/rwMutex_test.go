package lockex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type RWService struct {
	locker *sync.RWMutex
}

// 多寫互斥
// 寫讀互斥
// 多讀不會互斥
func Test_rwLock(t *testing.T) {
	svc := RWService{locker: &sync.RWMutex{}}
	count := 0
	times := 1000
	wg := new(sync.WaitGroup)
	wg.Add(times)
	write := func() {
		defer wg.Done()
		defer svc.locker.Unlock()
		svc.locker.Lock()
		count++
	}
	read := func() {
		defer svc.locker.RUnlock()
		svc.locker.RLock()
		fmt.Print(count, ",")
	}
	for i := 0; i < times; i++ {
		go write()
	}
	for i := 0; i < times; i++ {
		go read()
	}
	wg.Wait()
	fmt.Println(count)
}

func Test_rLock(t *testing.T) {
	svc := RWService{locker: &sync.RWMutex{}}
	count := 0
	times := 1000
	wg := new(sync.WaitGroup)
	wg.Add(times)
	call := func() {
		defer wg.Done()
		defer svc.locker.RUnlock()
		svc.locker.RLock()
		count++
	}
	for i := 0; i < times; i++ {
		go call()
	}
	wg.Wait()
	fmt.Println(count)
}

func Test_wLock(t *testing.T) {
	svc := RWService{locker: &sync.RWMutex{}}
	count := 0
	times := 1000
	wg := new(sync.WaitGroup)
	wg.Add(times)
	call := func() {
		defer wg.Done()
		defer svc.locker.Unlock()
		svc.locker.Lock()
		count++
	}
	for i := 0; i < times; i++ {
		go call()
	}
	wg.Wait()
	fmt.Println(count)
}

func Test_rLockSimple(t *testing.T) {
	m := new(sync.RWMutex)

	read := func(i int) {
		println(i, "read start")

		m.RLock()
		println(i, "reading")
		time.Sleep(1 * time.Second)
		m.RUnlock()

		println(i, "read over")
	}
	// 多个同时读
	go read(1)
	go read(2)

	time.Sleep(2 * time.Second)
}

func Test_wLockSimple(t *testing.T) {
	m := new(sync.RWMutex)
	write := func(i int) {
		println(i, "write start")

		m.Lock()
		println(i, "writing")
		time.Sleep(1 * time.Second)
		m.Unlock()

		println(i, "write over")
	}
	// 同時寫
	go write(1)
	go write(2)

	time.Sleep(2 * time.Second)
}
