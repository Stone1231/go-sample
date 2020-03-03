package goex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_runOnlyOnce(t *testing.T) {

	var do = func(o *sync.Once) {

		fmt.Println("Start do")

		//只會執行一次
		o.Do(func() {
			fmt.Println("Doing something...")
		})

		fmt.Println("Do end")
	}

	o := &sync.Once{}

	go do(o)

	go do(o)

	time.Sleep(time.Second * 2)
}
