package ioex

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_read(t *testing.T) {
	files := []string{"test.txt"}

	counts := make(map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Test_buffer(t *testing.T) {
	// var out io.Writer = os.Stdout // modified during
	// out = new(bytes.Buffer)
	var out io.Writer = new(bytes.Buffer)

	fmt.Fprint(out, "test test")

	got := out.(*bytes.Buffer).String()
	fmt.Println(got)
}
