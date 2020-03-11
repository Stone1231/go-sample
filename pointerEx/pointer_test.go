package pointerex

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_zerovalPointer(t *testing.T) {
	zeroptr := func(iptr *int) {
		fmt.Println("*iptr:", *iptr) // = i

		*iptr = 0

		fmt.Println("iptr:", iptr) // = &i
	}
	zeroval := func(ival int) {
		ival = 0

		fmt.Println("&ival:", &ival)
	}

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

func Test_mapLoop(t *testing.T) {
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

func Test_arrayLoop(t *testing.T) {
	m := [...]struct {
		name string
		age  int
	}{
		{"user1", 10},
		{"user2", 20},
		{"user3", 30},
	}
	for i, item := range m {
		fmt.Println(item) //item 只是一個指標指向不同的成員
		fmt.Printf("%p \n", &item)

		m[i].name = fmt.Sprintf("ptr_%v", i)
		item.name = fmt.Sprintf("val_%v", i) //無效

		v := m[i]
		v.name = fmt.Sprintf("val2_%v", i) //無效
		p := &m[i]
		p.name = fmt.Sprintf("ptr2_%v", i)
	}
	fmt.Println(m)
}

func Test_sliceLoop(t *testing.T) {
	m := []struct {
		name string
		age  int
	}{
		{"user1", 10},
		{"user2", 20},
		{"user3", 30},
	}
	for i, item := range m {
		fmt.Println(item) //同上, item 只是一個指標指向不同的成員
		fmt.Printf("%p \n", &item)

		m[i].name = fmt.Sprintf("ptr_%v", i)
		item.name = fmt.Sprintf("val_%v", i) //無效

		v := m[i]
		v.name = fmt.Sprintf("val2_%v", i) //無效
		p := &m[i]
		p.name = fmt.Sprintf("ptr2_%v", i)
	}
	fmt.Println(m)
}

func Test_sliceArrayAccess(t *testing.T) {
	arrayF := func(a [10]int) *[10]int {
		a[0] = 2
		return &a
	}
	sliceF := func(s []int) *[]int {
		s[0] = 2
		return &s
	}

	a := [10]int{}
	a[0] = 1
	s := make([]int, 10)
	s[0] = 1

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", arrayF(a))
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", sliceF(s))

	fmt.Println(fmt.Sprintf("%p", &a) == fmt.Sprintf("%p", arrayF(a)))
	fmt.Println(fmt.Sprintf("%p", &s) == fmt.Sprintf("%p", sliceF(s)))

	fmt.Println(a) //[1 0 0 0 0 0 0 0 0 0]
	fmt.Println(s) //[2 0 0 0 0 0 0 0 0 0]
}
