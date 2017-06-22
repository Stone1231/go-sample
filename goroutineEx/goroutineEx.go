package main

import (
	"math"
	"runtime"
	"sync"
)

func main() {
	//test1()
	//test12()
	//test2()
	test3()
}

///WaitGroup
func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}

func test1() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done() // = Add(-1)
			sum(id)
		}(i)
	}
	wg.Wait()
}

func test12() {
	for i := 0; i < 3; i++ {
		sum(i)
	}
}

///Goexit
func test2() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer println("A.defer")
		func() {
			defer println("B.defer")
			runtime.Goexit() // 终止当前 goroutine
			println("B")     // 不会执行
		}()
		println("A") // 不会执行
	}()
	wg.Wait()
}

///
func test3() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	var A = func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			println(i)
			if i == 3 {
				runtime.Gosched()
			}
		}
	}
	var B = func() {
		defer wg.Done()
		println("Hello, World!")
	}
	//先A後B
	go B()
	go A()

	wg.Wait()
}
