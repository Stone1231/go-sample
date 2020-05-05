package goex

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"testing"
)

func sum0toMax(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}

func Test_sum0toMax(t *testing.T) {
	for i := 0; i < 3; i++ {
		sum0toMax(i)
	}
}

//WaitGroup
func Test_sum0toMaxWaitGroup(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() // = Add(-1)
			sum0toMax(id)
		}(i)
	}
	wg.Wait()
}

///Goexit
func Test_goexitWaitGroup(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer println("A1.defer")
		defer println("A2.defer") //比A1先執行, Stack
		func() {
			defer println("B.defer")
			runtime.Goexit() // 终止当前 goroutine
			println("B")     // 不会执行
		}()
		println("A") // 不会执行
	}()
	wg.Wait()
}

func Test_multiGoWaitGroup(t *testing.T) {
	//func multiGo() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	var A = func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			println("A", i)
			//time.Sleep(1 * time.Second)
		}
	}

	var B = func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			println("B", i)
			//time.Sleep(1 * time.Second)
		}
	}

	var C = func() {
		defer wg.Done()
		//go B()
		for i := 0; i < 100; i++ {
			println("C", i)
			//time.Sleep(1 * time.Second)
		}
	}

	//先A後B?

	go C()
	go B()
	go A()

	wg.Wait()
}

//WaitGroup
func Test_WaitGroup(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(" ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
