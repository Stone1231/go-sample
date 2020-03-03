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
	type data struct {
		a int
		b string
	}
	var d = data{1234, "abcd"}
	var p *data
	p = &d

	fmt.Printf("%p, %v, %v %v, %v\n", p, p.a, d.a, p.b, d.b)
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

func Test_MapLoop(t *testing.T) {
	m := map[int]struct {
		name string
		age  int
	}{
		1: {"user1", 10},
		2: {"user2", 20},
		3: {"user3", 30},
	}
	m2 := map[string]*struct {
		name string
		age  int
	}{}

	for _, item := range m {
		m2[item.name] = &item
	}

	for _, item := range m2 {
		fmt.Println(item)
	}
	// &{user3 30}
	// &{user3 30}
	// &{user3 30}
}

func Test_ArrayLoop(t *testing.T) {
	m := [...]struct {
		name string
		age  int
	}{
		{"user1", 10},
		{"user2", 20},
		{"user3", 30},
	}
	for _, item := range m {
		fmt.Println(item) //複製品
		fmt.Printf("%p \n", &item)
	}
}
