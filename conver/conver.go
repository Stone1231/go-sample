package main

import (
	"fmt"
)

func main() {
	test1()
}

func test1() {
	var b byte = 100
	// var n int = b // Error: cannot use b (type byte) as type int in assignment
	var n int = int(b) //
	fmt.Println(n)
}
