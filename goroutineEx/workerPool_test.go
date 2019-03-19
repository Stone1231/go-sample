package goex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_worker(t *testing.T) {

	jobs := make(chan int, 3)

	results := make(chan int, 3)

	var wg sync.WaitGroup

	worker := func(id int) {
		for j := range jobs {
			fmt.Println("worker", id, "processing job", j)

			time.Sleep(time.Second)
			results <- j * 2
		}
		wg.Done()
	}

	n := 3 // tree works
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go worker(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for r := range results {
		fmt.Println(r)
	}
}
