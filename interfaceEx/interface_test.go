package interfaceex

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func Test_pointerOrValue(t *testing.T) {
	u := &User{1, "Stone"} //User{1, "Stone"} 一樣
	u.UpdateNameValue("Stone Value")
	fmt.Println(u)
	u.UpdateNamePointer("Stone Pointer")
	fmt.Println(u)
}

func Test_printerInterface(t *testing.T) {
	var p Printer = &User{1, "Tom"} // *User 方法集包含 String、Print。
	p.Print()
}

func Test_refInterface(t *testing.T) {
	u := User{1, "Stone"}
	var i interface{} = u  //複製品
	var j = u              //複製品
	var k = &u             //變更會互相影響
	var l interface{} = &u //變更會互相影響

	u.id = 11
	u.name = "Stone1"

	fmt.Printf("u:%v\n", u)
	fmt.Printf("j:%v\n", j)

	j.id = 2
	j.name = "Lisa"

	fmt.Printf("u:%v\n", u)

	k.id = 3
	k.name = "Rita"

	fmt.Printf("u:%v\n", u)
	fmt.Printf("i:%v\n", i.(User))
	fmt.Printf("j:%v\n", j)

	l.(*User).id = 4
	l.(*User).name = "Ava"
	fmt.Printf("u:%v\n", u)
}

func Test_emptyInterface(t *testing.T) {
	PrintAny(1)
	PrintAny("Hello, World!")
}

func Test_inheriance(t *testing.T) {
	m := Manager{User{1, "Tom"}}
	fmt.Println(m.String())
}

func Test_receiver(t *testing.T) {
	u := User{1, "Tom"}
	mFun := u.Print
	mFun()
	u.id, u.name = 2, "Jack"
	mFun()

	mFun2 := (*User).Print
	mFun2(&u)
}

func Test_attrInterface(t *testing.T) {
	o := Attr{&User{1, "Tom"}}
	fmt.Println(o.s.String())
}

func Test_uintptrUnsafe(t *testing.T) {
	// tab & data都是nil時,介面才是nil
	var a interface{} = nil         // tab = nil, data = nil
	var b interface{} = (*int)(nil) // tab包含*int, data = nil
	type iface struct {
		itab, data uintptr
	}
	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	// tab = nil, data = nil
	// tab    *int     , data = nil
	fmt.Println(a == nil, ia)
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())
}

func Test_converInterface(t *testing.T) {
	var o Printer = &User{1, "Tom"}
	var s Stringer = o
	fmt.Println(s.String())
}

func Test_converPointer(t *testing.T) {
	var o interface{} = &User{1, "Stone1"}
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println("fmt.Stringer")
		fmt.Println(i)
	}
	if i, ok := o.(Stringer); ok {
		fmt.Println("Stringer")
		fmt.Println(i)
	}
	u := o.(*User)
	//u := o.(User)  error
	fmt.Println("u")
	fmt.Println(u)

	fmt.Println("switch")
	var o2 interface{} = &User{2, "Stone2"}
	switch v := o2.(type) {
	case nil:
		fmt.Println("nil")
	case fmt.Stringer:
		fmt.Println(v)
	case func() string:
		fmt.Println(v())
	case *User:
		fmt.Printf("%d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown")
	}
}

func Test_funInterface(t *testing.T) {

	type iDo interface {
		Do()
	}

	var o iDo = FuncDo(
		func() {
			println("Hello, World!")
		})
	o.Do()
}
