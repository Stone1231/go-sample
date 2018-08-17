package pointerex

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_zerovalPointer(t *testing.T) {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("pointer:", &i)

}

func Test_structPointer(t *testing.T) {
	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d

	fmt.Printf("%p, %v, %v\n", p, p.a, d.a)
}

func Test_refScope(t *testing.T) {
	s := "abc"
	fmt.Println(&s)
	s, y := "hello", 20
	fmt.Println(&s, y)

	{
		s, z := 1000, 30
		fmt.Println(&s, z)
	}

}

func Test_byteArray(t *testing.T) {
	x := 0x12345678
	p := unsafe.Pointer(&x)
	n := (*[4]byte)(p)
	for i := 0; i < len(n); i++ {
		fmt.Printf("%x ", n[i])
	}
}

func Test_uintptrOffset(t *testing.T) {
	d := struct {
		s string
		x int
	}{"abc", 100}

	p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(d.x)        // uintptr + offset

	px := (*int)(unsafe.Pointer(p)) // Pointer -> *int
	*px = 200                       // d.x = 200

	fmt.Printf("%#v\n", d)

	p = uintptr(unsafe.Pointer(&d))
	p += unsafe.Offsetof(d.s)
	ps := (*string)(unsafe.Pointer(p))
	*ps = "def"

	fmt.Printf("%#v\n", d)
}
