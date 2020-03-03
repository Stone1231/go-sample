package goex

import (
	"fmt"
	"os"
	"testing"
)

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

func Test_createConsumer(t *testing.T) {
	data := createConsumer()
	data <- 1
	data <- 2
	close(data)

	//沒有可用channel, 阻塞main goroutine
	//讓main函數不退出，讓它在後台一直執行
	select {}
}
