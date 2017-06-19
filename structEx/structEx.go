package main

import (
	"fmt"
)

func main() {

	// test1()

	// test2()

	// test3()

	//testMap()

	//test5()

	//test7()

	//test8()

	//test9()

	//test10()

	test11()
}
func test1() {
	n1 := Node{
		id:   1,
		data: nil,
	}
	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}
	fmt.Println(n2.id)
}

type Node struct {
	_    int
	id   int
	data *byte
	next *Node
}

//
func test2() {
	type User struct {
		name string
		age  int
	}
	u1 := User{"Tom", 20}
	u1.age = 18
	//u2 := User{"Tom"} //error 初始化全部成員都要設才行
}

//
func test3() {
	type File struct {
		name string
		size int
		attr struct {
			perm  int
			owner int
		}
	}
	f := File{
		name: "test.txt",
		size: 1025,
		// attr: {0755, 1}, // Error: missing type in composite literal
	}

	f.attr.owner = 1
	f.attr.perm = 0755
	var attr = struct {
		perm  int
		owner int
	}{2, 0756} //0...8進位, 0x...16進位

	f.attr = attr

	fmt.Println(f.attr.owner)
}

//
func testMap() {
	type User struct {
		id   int
		name string
	}
	m := map[User]int{ //User可做map的key
		User{1, "Tom"}: 100,
	}

	a := User{1, "Tom"}

	u := m[a]

	fmt.Println(u)

	//fmt.Println(&a)
	fmt.Printf("%p\n", &a)

	for k := range m {
		//fmt.Println(&k)
		fmt.Printf("%p\n", &k)
	}
}

//
func test5() {
	var u1 struct {
		name string "username"
	}
	var u2 struct{ name string }
	//u2 = u1 // Error: cannot use u1 (type struct { name string "username" }) as
	//        type struct { name string } in assignment
	u1.name = "name1"
	fmt.Println(u1.name)
	u2.name = "name2"
}

//
func testNull() {
	var null struct{}
	set := make(map[string]struct{})
	set["a"] = null
}

//
func test7() {
	type User struct {
		name string
	}
	type Manager struct {
		User
		title string
	}

	m := Manager{
		User:  User{"Tom"}, //
		title: "Administrator",
	}
	fmt.Println(m.name, m.User.name)
}

//
func test8() {
	type Resource struct {
		id int
	}
	type User struct {
		Resource
		name string
	}
	type Manager struct {
		User
		title string
	}
	var m Manager
	m.id = 1
	m.name = "Jack"
	m.title = "Administrator"
	fmt.Println(m)
}

//
func test9() {

	type Resource struct {
		id   int
		name string
	}
	type Classify struct {
		id int
	}
	type User struct {
		Resource
		Classify
		name string
	}
	u := User{
		Resource{1, "people"},
		Classify{100},
		"Jack",
	}
	println(u.name)
	println(u.Resource.name)
	// println(u.id)// Error: ambiguous selector u.id 欄位重覆
	println(u.Classify.id)
}

//
func test10() {
	type Resource struct {
		id int
	}
	type User struct {
		*Resource
		//Resource
		name string
	}
	u := User{
		&Resource{1},
		//Resource{1},
		"Administrator",
	}
	println(u.id)
	println(u.Resource.id)
}

//
func test11() {
	type User struct {
		id   int
		name string
	}
	type Manager struct {
		User
		title string
	}

	m := Manager{User{1, "Tom"}, "Administrator"}
	// var u User = m // Error: cannot use m (type Manager) as type User in assignment //  互休       互
	var u User = m.User //
	fmt.Println(u.name)
}
