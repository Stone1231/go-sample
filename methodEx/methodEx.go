package main

import "fmt"

func main() {
	//test1()

	//test2()

	//test3()

	//test4()

	//test5()

	//test6()

	//test7()

	test8()
}

///
type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{make([]interface{}, 10)}
}
func (*Queue) Push(e interface{}) error {
	panic("not implemented")
}

// func (Queue) Push(e int) error {
//     panic("not implemented")
// }
func (self *Queue) length() int {
	return len(self.elements)
}

/// receiver T 和 *T 的差别
type Data struct {
	x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &self)
}
func (self *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", self)
}
func test1() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()   // ValueTest(d)
	d.PointerTest() // PointerTest(&d)
	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
}

/// 不支援pointer查找方法成員
type X struct{}

func (*X) test() {
	println("X.test")
}
func test2() {
	p := &X{}
	p.test()
	// Error: calling method with receiver &p (type **X) requires explicit dereference
	// (&p).test()
}

///
type User struct {
	id   int
	name string
}
type Manager struct {
	User
}

func (self *User) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %v", self, self)
}
func test3() {
	m := Manager{User{1, "Tom"}}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
}

///透過匿名字段 可以獲得和繼承類似的復用能力,
///在外層定義同名方法就可以實現override
type User2 struct {
	id   int
	name string
}
type Manager2 struct {
	User2
	title string
}

func (self *User2) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}
func (self *Manager2) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}
func test4() {
	m := Manager2{User2{1, "Tom"}, "Administrator"}
	fmt.Println(m.ToString())
	fmt.Println(m.User2.ToString())
}

///
type User3 struct {
	id   int
	name string
}

func (self *User3) test() {
	fmt.Printf("%p, %v\n", self, self)
}

func test5() {
	u := User3{1, "Tom"}
	u.test()
	mValue := u.test
	mValue() //隐式傳遞 receiver
	mExpression := (*User3).test
	mExpression(&u) // 顯式傳遞 receiver
}

///
type User4 struct {
	id   int
	name string
}

func (self User4) test() {
	fmt.Println(self)
}
func test6() {
	u := User4{1, "Tom"}
	mValue := u.test // ⽴立即复制 receiver，因为不是指针类型，不受后续修改影响。
	u.id, u.name = 2, "Jack"
	u.test()
	mValue()
}

///
type User5 struct {
	id   int
	name string
}

func (self *User5) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}
func (self User5) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

func test7() {
	u := User5{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &u, u)
	mv := User5.TestValue
	mv(u)
	mp := (*User5).TestPointer
	mp(&u)
	mp2 := (*User5).TestValue // *User 方法集包含 TestValue。
	mp2(&u)                   // 签名变为 func TestValue(self *User)。
	//                           实际依然是 receiver value copy。
}

///将方法 "还原" 成函数，就容易理解
type Data2 struct{}

func (Data2) TestValue()    {}
func (*Data2) TestPointer() {}
func test8() {
	var p *Data2 = nil
	p.TestPointer()
	(*Data2)(nil).TestPointer() // method value
	(*Data2).TestPointer(nil)   // method expression
	//p.TestValue()               // invalid memory address or nil pointer dereference
	// (Data)(nil).TestValue()     // cannot convert nil to type Data
	// Data.TestValue(nil)         // cannot use nil as type Data in function argument
}
