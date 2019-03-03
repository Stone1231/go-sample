package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var count int

func sumLenHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["txt"]
	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	fmt.Println("Url Param 'key' is: " + string(key))

	length := len(key)

	count += length
	fmt.Fprintf(w, "total strings length: %d\n", count)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	count = 0
	w.Write([]byte("OK\n"))
}

var limiter = rate.NewLimiter(30, 5)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func startAPI() {
	//http.HandleFunc("/", sumLenHandler)
	//log.Fatal(http.ListenAndServe("localhost:"+API_PORT, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("/", sumLenHandler)
	mux.HandleFunc("/reset", resetHandler)

	log.Fatal(http.ListenAndServe("localhost:"+API_PORT, limit(mux)))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	if string(body) != "Too Many Requests\n" {
		ch <- fmt.Sprintf("%.2fs %s", secs, body)
	} else {
		close(ch)
	}
}
