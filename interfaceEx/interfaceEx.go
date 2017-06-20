package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//test1()

	//test2()

	//test3()

	//test4()

	//test5()

	test6()
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
func test1() {
	var t Printer = &User{1, "Tom"} // *User 方法集包含 String、Print。
	t.Print()
}

///
func Print(v interface{}) { //無任何方法,所有的類別都實現了空介面
	fmt.Printf("%T: %v\n", v, v)
}
func test2() {
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
func test3() {
	t := Tester{&User2{1, "Tom"}}
	fmt.Println(t.s.String())
}

///
type User3 struct {
	id   int
	name string
}

func test4() {
	u := User3{1, "Tom"}
	var i interface{} = u //複製品
	u.id = 2
	u.name = "Jack"
	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", i.(User3))
}

///介面轉型 只有指標才能修改狀態
type User4 struct {
	id   int
	name string
}

func test5() {
	u := User4{1, "Tom"}
	var vi, pi interface{} = u, &u
	//vi.(User).name = "Jack" // Error: cannot assign to vi.(User).name
	pi.(*User4).name = "Jack"
	fmt.Printf("%v\n", vi.(User4))
	fmt.Printf("%v\n", pi.(*User4))
}

///
func test6() { // tab ,data都是nil,介面才是nil
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
