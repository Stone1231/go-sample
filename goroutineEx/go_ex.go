package goex

import (
	"fmt"
	"math"
	"os"
)

func sum0toMax(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}

//簡單工廠模式
func createConsumer() chan int {
	data := make(chan int, 3)
	go func() {
		for d := range data {
			fmt.Println(d)
		}
		os.Exit(0)
	}()
	return data
}
