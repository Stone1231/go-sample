package main

import (
	"fmt"
	"unsafe"
)

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d

	fmt.Printf("%p, %v\n", p, p.a)
}

func test2() {
	x := 0x12345678
	p := unsafe.Pointer(&x)
	n := (*[4]byte)(p)
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}
}

func test3() {
	d := struct {
		s string
		x int
	}{"abc", 100}
	p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(d.x)        // uintptr + offset

	p2 := unsafe.Pointer(p) // uintptr -> Pointer
	px := (*int)(p2)        // Pointer -> *int
	*px = 200               // d.x = 200
	fmt.Printf("%#v\n", d)

}
