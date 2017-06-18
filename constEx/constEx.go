package main

import (
	"fmt"
	"unsafe"
)

const x, y int = 1, 2
const s = "Hello, World!"

// 多常量初始化 // 类型推断
// 常量组
const (
	a, b      = 10, 100
	c    bool = false
)

const (
	s2 = "abc"
	x2 // x2 = "abc"
)

const (
	a2 = "abc"
	b2 = len(a2)
	c2 = unsafe.Sizeof(b)
)

const (
	a3 byte = 100 // int to byte
	//b3 int  = 1e20 // float64 to int, overflows
)

const (
	Sunday    = iota //0
	Monday           // 1,通常省略后续 表达式。
	Tuesday          //2
	Wednesday        //3
	Thursday         //4
	Friday           //5
	Saturday         //6
)

const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) //iota=1
	MB                          //與KB表達相同, iota = 2
	GB
	TB
)

func main() {
	const x = "xxx" // 未使 局部常量不会引发编译错误。
	// fmt.Print(x2)
	// fmt.Print(Saturday)
	// fmt.Print(MB)

	fmt.Print(17 & 3)
}
