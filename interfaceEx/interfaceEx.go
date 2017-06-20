package main

import "fmt"

func main() {
	test1()

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
