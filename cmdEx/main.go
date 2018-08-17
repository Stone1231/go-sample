package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	hello(args)
	// linux  $go run *.go
	//windows 1. go build 2. run exe
}
