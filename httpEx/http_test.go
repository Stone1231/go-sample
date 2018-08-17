package httpex

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

//http.Get(url)  chan
func Test_fetchall(t *testing.T) {
	urls := []string{"http://www.ncu.edu"}
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func Test_server(t *testing.T) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
