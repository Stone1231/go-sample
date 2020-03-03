package goex

import (
	"fmt"
	"testing"
	"time"
)

func Test_onlySendReceive(t *testing.T) {
	c := make(chan int, 3)
	var send chan<- int = c // send-only
	var recv <-chan int = c // receive-only

	//單純傳送接收
	send <- 1
	<-recv
	// <-send // Error: receive from send-only type chan<- int
	// recv <- 2 // Error: send to receive-only type <-chan int

	go func() {
		time.Sleep(time.Second * 2)
		send <- 1
	}()

	n := <-recv //等到接收為止
	fmt.Println(n)
}
