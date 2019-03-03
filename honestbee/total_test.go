package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestApi(t *testing.T) {
	startAPI()
}

func TestTcp(t *testing.T) {
	startTCP()
}

func TestFetchall(t *testing.T) {

	go startAPI()

	//urls := []string{"http://www.ncu.edu"}
	txt := url.QueryEscape("!@#$%^&*()")
	encodeurl := "http://localhost:" + API_PORT + "/?txt=" + txt
	urls := []string{encodeurl}

	//start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func TestStart(t *testing.T) {
	go startAPI()

	startTCP()
}

func BenchmarkFetchall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//txt := url.QueryEscape("!@#$%^&*()")
		txt := url.QueryEscape("1")
		encodeurl := "http://localhost:" + API_PORT + "/?txt=" + txt
		urls := []string{encodeurl}

		//start := time.Now()
		ch := make(chan string)
		for _, url := range urls {
			go fetch(url, ch) // start a goroutine
		}
		for range urls {
			res := <-ch
			if res != "" {
				fmt.Println(res)
			}
		}
	}
}
