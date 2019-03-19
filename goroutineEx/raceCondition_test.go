package goex

import (
	"fmt"
	"testing"
)

func Test_race(t *testing.T) {
	a, b := 0, 0
	times := 10000
	c := make(chan bool)

	for i := 0; i < times; i++ {
		go func() {
			a++
			c <- true
		}()
	}

	for i := 0; i < times; i++ {
		<-c
		b++
	}
	fmt.Printf("a = %d\n", a) //可能會得到9xxx
	fmt.Printf("b = %d\n", b)
}
