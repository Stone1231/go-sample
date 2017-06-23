package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
)

func main() {
	//test1()
	//test12()
	//test2()
	//test3()
	//test32()
	//test4()
	//test5()
	//test6()
	//test7()
	test8()
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

func test32() {
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
		go B()
		for i := 0; i < 100; i++ {
			println("C", i)
			//time.Sleep(1 * time.Second)
		}
	}

	println("gogogo")
	//先A後B

	go C()
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

//有問題？
func test5() {
	data := make(chan int, 3) //緩衝區可儲存3個元素
	exit := make(chan bool)
	data <- 1 //緩衝區未滿前不會阻塞
	data <- 2
	data <- 3
	//data <- 4 // all goroutines are asleep - deadlock!//如果緩衝區已滿,阻塞
	go func() {
		for d := range data { // 緩衝區為空前不會阻塞
			fmt.Println(d)
		}
		exit <- true
	}()
	data <- 4
	data <- 5
	for index := 0; index < 3000; index++ {
		data <- index
	}
	fmt.Println("xxxx")
	close(data)
	fmt.Println("vvv")
	<-exit
	fmt.Println("ooo")
}

//
func test6() {
	//ok-idiom 模式判斷channel是否關閉
	// for {
	// 	if d, ok := <-data; ok {
	// 		fmt.Println(d)
	// 	} else {
	// 		break
	// 	}
	// }

	d1 := make(chan int)
	d2 := make(chan int, 3)
	d2 <- 1
	fmt.Println(len(d1), cap(d1)) // 0  0
	fmt.Println(len(d2), cap(d2)) // 1  3
}

//
func test7() {
	c := make(chan int, 3)
	var send chan<- int = c // send-only
	var recv <-chan int = c // receive-only

	send <- 1
	// <-send // Error: receive from send-only type chan<- int

	<-recv
	// recv <- 2 // Error: send to receive-only type <-chan int
}

//select channel
func test8() {
	a, b := make(chan int, 3), make(chan int)
	go func() {
		v, ok, s := 0, false, ""
		for {
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}
			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()
	for i := 0; i < 5; i++ {
		select {
		case a <- i:
		case b <- i:
		}
	}
	close(a)
	select {} //沒有可用channel, 阻塞main goroutine
}
