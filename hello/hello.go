package main

import (
	"fmt"
	"math"
)

var x, y, z int
var s, n = "abc", 123
var (
	a int
	b float32
)

func main() {
	n, s := 0x1234, "Hello, World!"
	fmt.Println(x, s, n)

	_, s2 := test()
	fmt.Println(s2)

	ref()

	mathFun()
}

func test() (int, string) {
	return 1, "abc"
}

func ref() {
	s := "abc"
	fmt.Println(&s)
	s, y := "hello", 20
	fmt.Println(&s, y)

	{
		s, z := 1000, 30
		fmt.Println(&s, z)
	}
	// 0xc42006e1a0
	// 0xc42006e1a0 20
	// 0xc42006e1c0 30
}
func mathFun() {
	a1, b1, c1, d1 := 071, 0x1F, 1e9, math.MinInt16
	fmt.Println(a1, b1, c1, d1)
}
