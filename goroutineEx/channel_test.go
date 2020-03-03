package goex

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

func Test_rangeChan(t *testing.T) {
	data := make(chan int)  //數據交換
	exit := make(chan bool) //退出通知

	go func() {
		for d := range data { //佇列接收通知直到close
			time.Sleep(time.Second * 2)
			fmt.Println(d)
		}
		fmt.Println("recv over.")
		exit <- true //發出退出通知(用false也一樣 )
	}()

	data <- 1 //發出數據
	data <- 2
	data <- 3

	close(data) //關閉佇列

	fmt.Println("send over.")

	<-exit //等待退出通知
	fmt.Println("exit")
}

func Test_rangeChanBreakLoop(t *testing.T) {
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

		// 同上
		// for {
		// 	if d, ok := <-data; ok {
		// 		fmt.Println(d)
		// 	} else {
		// 		break
		// 	}
		// }

		exit <- true
	}()

	data <- 4
	data <- 5

	end := make(chan bool)
	//end := false
	go func() {
		defer fmt.Println("after close")
		defer func() {
			end <- true
			close(data)

			//不好的寫法
			// end = true
			// time.Sleep(time.Millisecond * 1) //不停一下,持續data <- index會當掉
			// close(data)
		}()
		defer fmt.Println("before close")
		time.Sleep(time.Millisecond * 10)
	}()

L1:
	for index := 6; true; index++ {
		select {
		case data <- index:
		case <-end:
			break L1
		default: //若沒有這個判斷,data <- index會等到chan有空再執行
			fmt.Println(index, " channel is full !")
			//可以回復 服務繁忙，請稍微再試。
		}

		// select {
		// case <-end:
		// 	break L1
		// default:
		// 	data <- index
		// }

		// if end {
		// 	break
		// } else {
		// 	data <- index
		// }
	}

	<-exit
	fmt.Println("exit")
}

func Test_closeFunsChan(t *testing.T) {
	done := make(chan struct{})

	fmt.Printf("start\n")

	go func() {
		//os.Stdin.Read(make([]byte, 1)) // read a single byte
		time.Sleep(time.Second * 2)
		close(done)
	}()

	go func() {
		<-done
		fmt.Printf("done-1\n")
	}()

	go func() {
		<-done
		fmt.Printf("done-2\n")
	}()

	go func() {
		<-done
		fmt.Printf("done-3\n")
	}()

	go func() {
		<-done
		fmt.Printf("done-4\n")
	}()

loop:
	for {
		select {
		case <-done:
			time.Sleep(time.Second * 1)
			break loop
		}
	}

	fmt.Printf("end\n")
}

//select 實現超時
func Test_timeoutBySelect(t *testing.T) {
	w := make(chan bool)
	c := make(chan int, 2)
	go func() {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(time.Second * 2):
			fmt.Println("timeout 2.")
		}
		w <- true
	}()
	//c <- 1 // 註解掉引發timeout
	<-w
}

//使用closed channel發出退出通知
func Test_quitByClosedChannel(t *testing.T) {
	var wg sync.WaitGroup
	quit := make(chan bool)

	f := func(id int) {
		defer wg.Done()
		task := func() {
			println(id, time.Now().Nanosecond())
			time.Sleep(time.Second)
		}
		for {
			select {
			case <-quit: //closed channel不會阻塞,可用作退出通知
				return
			default: //執行正常任務
				task()
			}
		}
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go f(i)
	}

	time.Sleep(time.Second * 5) //讓goroutine執行一段時間
	close(quit)                 //發出退出通知
	wg.Wait()
	println("the end")
}

///channel 實現號誌(Semaphore)
//Semaphore是一件可以容納N人的房間，如果人不滿就可以進去，
//如果人滿了，就要等待有人出來。
//對於N=1的情況，稱為binary semaphore。
//一般的用法是，用於限制對於某一資源的同時訪問。
func Test_semaphore(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)
	sem := make(chan int, 1)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			sem <- 1 //發送給sem, 阻塞或成功
			for x := 0; x < 3; x++ {
				fmt.Println(id, x)
			}
			<-sem //接收數據後,使其他阻塞可以發送數據
		}(i)
	}
	wg.Wait()
}

func Test_waitGoFinish(t *testing.T) {
	c := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("B")
		c <- 1
	}()

	fmt.Println("A")

	<-c

	fmt.Println("C")
}

func Test_lenCap(t *testing.T) {
	d1 := make(chan int)
	d2 := make(chan int, 3)
	// d1 <- 2 //無緩衝又沒接收,會error
	d2 <- 1
	fmt.Println(len(d1), cap(d1)) // 0  0
	fmt.Println(len(d2), cap(d2)) // 1  3
}

func Test_selectChannel(t *testing.T) {
	a, b := make(chan int, 3), make(chan int)
	go func() {
		v, ok, s := 0, false, ""
		for {
			select {
			case v, ok = <-a: //加上ok布林值表示是否成功接收
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

	for i := 0; i < 30; i++ {
		select {
		case a <- i:
		case b <- i:
		}
	}

	close(a)

	//不需要
	//select {}
}
