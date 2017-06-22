package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {
	//test1()
	//test12()
	//test2()
	//test3()
	//test4()
	test5()
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

///
func test4() {
	data := make(chan int)  //數據交換
	exit := make(chan bool) //退出通知

	go func() {
		for d := range data { //佇列接收通知直到close
			fmt.Println(d)
		}
		fmt.Println("recv over.")
		exit <- true //發出退出通知
	}()

	data <- 1 //發出數據
	data <- 2
	data <- 3

	close(data) //關閉佇列

	fmt.Println("send over.")

	<-exit //等待退出通知
}

//
func test5() {
	data := make(chan int, 3) //緩衝區可儲存3個元素
	exit := make(chan bool)
	data <- 1 //緩衝區未滿前不會阻塞
	data <- 2
	data <- 3
	go func() {
		for d := range data { // 緩衝區為空前不會阻塞
			fmt.Println(d)
		}
		exit <- true
	}()
	//data <- 4 //如果緩衝區已滿,阻塞
	//data <- 5
	<-exit

}
