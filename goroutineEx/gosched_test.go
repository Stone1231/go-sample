package goex

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_gosched(t *testing.T) {
	say := func(s string) {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
			runtime.Gosched() //讓出CPU時間片。就像跑接力賽
		}
	}

	go say("world")

	go say("hello")
}
