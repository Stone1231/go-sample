package goex

import (
	"fmt"
	"testing"
	"time"
)

func Test_simple(t *testing.T) {
	//func printGo() {
	go func() {
		for {
			fmt.Println("A")
		}
	}()

	for {
		fmt.Println("B")
	}
}

func Test_twoAssync(t *testing.T) {
	var x, y int
	go func() {
		x = 1                   //A1
		fmt.Print("y:", y, " ") //A2
	}()

	go func() {
		y = 2                   //B1
		fmt.Print("x:", x, " ") //B2
	}()

	//A1,B1,A2,B2 OR B1,A1,A2,B2..各種順序都有可能

	time.Sleep(time.Millisecond * 1)
}

func Test_multiProcess(t *testing.T) {

	type Request struct {
		data []int
		ret  chan int
	}

	Process := func(req *Request) {
		x := 0
		for _, i := range req.data {
			x += i
		}
		time.Sleep(2 * time.Second)
		req.ret <- x
	}

	req := &Request{[]int{10, 20, 30}, make(chan int, 1)}
	req2 := &Request{[]int{11, 22, 33}, make(chan int, 1)}

	go Process(req)
	go Process(req2) //go others...

	fmt.Println(<-req.ret)
	fmt.Println(<-req2.ret)
}
