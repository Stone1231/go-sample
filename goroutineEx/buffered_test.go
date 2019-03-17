package goex

import (
	"fmt"
	"testing"
	"time"
)

func Test_buffered(t *testing.T) {

	chJob := func(title string, ch chan int, count int) {

		s := ""
		go func() {
			for i := 0; i < count; i++ {
				ch <- i
				s += "."
			}
			close(ch)
			s += " close "
		}()

		for i := range ch {
			s += fmt.Sprintf("%d", i)
		}

		fmt.Println(title + s)
	}

	go chJob("  : ", make(chan int), 10)     //unbuffered
	go chJob(" 0: ", make(chan int, 0), 10)  //unbuffered
	go chJob(" 1: ", make(chan int, 1), 10)  //buffered
	go chJob(" 3: ", make(chan int, 3), 10)  //buffered
	go chJob("10: ", make(chan int, 10), 10) //buffered

	<-time.After(time.Second * 1)
}

func Test_unbuffered(t *testing.T) {
	data := make(chan int)

	go func() {
		for i := range data { //佇列接收通知直到close
			fmt.Println(i)
		}

		fmt.Println("the end") //close 後才會執行到
	}()

	data <- 1
	data <- 2
	data <- 3
	data <- 4

	close(data)

	<-time.After(time.Second * 1)
}
