package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//printerInterface()

	//smartyInterface()

	//attrInterface()

	//printerInterface2()

	//printerInterface3()

	//uintptrUnsafe()

	//converPointer()

	//converPointer2()

	methodInterface()
}

type Stringer interface {
	String() string
}
type Printer interface {
	Stringer // 接口嵌入。
	Print()
}
type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}
func (self *User) Print() {
	fmt.Println(self.String())
}
func printerInterface() {
	var t Printer = &User{1, "Tom"} // *User 方法集包含 String、Print。
	t.Print()
}

///
func Print(v interface{}) { //無任何方法,所有的類別都實現了空介面
	fmt.Printf("%T: %v\n", v, v)
}
func smartyInterface() {
	Print(1)
	Print("Hello, World!")
}

///
type Tester struct {
	s interface {
		String() string
	}
}

type User2 struct {
	id   int
	name string
}

func (self *User2) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}
func attrInterface() {
	t := Tester{&User2{1, "Tom"}}
	fmt.Println(t.s.String())
}

///
type User3 struct {
	id   int
	name string
}

func printerInterface2() {
	u := User3{1, "Tom"}
	var i interface{} = u //複製品
	var j = u             //複製品
	var k = &u            //變更會互相影響

	u.id = 2
	u.name = "Jack"

	fmt.Printf("u:%v\n", u)
	fmt.Printf("j:%v\n", j)

	j.id = 3
	j.name = "lisa"

	fmt.Printf("u:%v\n", u)

	k.id = 4
	k.name = "rita"

	fmt.Printf("u:%v\n", u)
	fmt.Printf("i:%v\n", i.(User3))
	fmt.Printf("j:%v\n", j)

}

///介面轉型 只有指標才能修改狀態
type User4 struct {
	id   int
	name string
}

func printerInterface3() {
	u := User4{1, "Tom"}
	var vi, pi interface{} = u, &u
	//vi.(User).name = "Jack" // Error: cannot assign to vi.(User).name
	pi.(*User4).name = "Jack"
	fmt.Printf("%v\n", vi.(User4))
	fmt.Printf("%v\n", pi.(*User4))
}

///
func uintptrUnsafe() { // tab ,data都是nil,介面才是nil
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

///
type User5 struct {
	id   int
	name string
}

func (self *User5) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}

// func (self User5) String() string {
// 	return fmt.Sprintf("%d, %s", self.id, self.name)
// }

func converPointer() {
	var o interface{} = &User5{1, "Tom"}
	if i, ok := o.(fmt.Stringer); ok { // ok-idiom
		fmt.Println(i)

	}
	u := o.(*User5)
	// u := o.(User)
	fmt.Println(u)

	// var o2 interface{} = User5{1, "Tom"}
	// if i, ok := o2.(fmt.Stringer); ok { // ok-idiom
	// 	fmt.Println(i)

	// }
	// u2 := o2.(User5)
	// // u := o.(User5)
	// fmt.Println(u2)
	fmt.Println("switch")
	var o3 interface{} = &User5{1, "Tom"}
	switch v := o3.(type) {
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

///
type Stringer2 interface {
	String() string
}
type Printer2 interface {
	String() string
	Print()
}

type User6 struct {
	id   int
	name string
}

func (self *User6) String() string {
	return fmt.Sprintf("%d, %v", self.id, self.name)
}
func (self *User6) Print() {
	fmt.Println(self.String())
}
func converPointer2() {
	var o Printer2 = &User6{1, "Tom"}
	var s Stringer2 = o
	fmt.Println(s.String())
}

///
//var _ fmt.Stringer = (*Data)(nil)
type Tester2 interface {
	Do()
}
type FuncDo func()

func (self FuncDo) Do() { self() }
func methodInterface() {
	var t Tester2 = FuncDo(func() { println("Hello, World!") })
	t.Do()
}
