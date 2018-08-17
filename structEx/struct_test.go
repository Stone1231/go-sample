package structex

import (
	"fmt"
	"log"
	"testing"
	"unsafe"
)

// func main() {

// nodeData()

// newInit()

// attrStruct()

// mapKeyUseStruct()

// attrOtherName()

// attrStruct2()

// attrStruct3()

// attrStruct4()

// pointerStruct()

// 	newInit2()
// }

type Node struct {
	_    int
	id   int
	data *byte
	next *Node
}

func Test_node(t *testing.T) {
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
	fmt.Println(n2.next.id)
}

//
func Test_attr(t *testing.T) {
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
		attr: struct {
			perm  int
			owner int
		}{4, 3},
	}
	fmt.Println(f.attr.owner)

	f.attr.owner = 1
	f.attr.perm = 0755
	fmt.Println(f.attr.owner)

	var attr = struct {
		perm  int
		owner int
	}{2, 0756} //0...8進位, 0x...16進位
	f.attr = attr
	fmt.Println(f.attr.owner)
}

func Test_attr2(t *testing.T) {
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
	fmt.Println(m.name, m.User.name, m.title)
}

func Test_attrDuplicated(t *testing.T) {

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

func Test_inheriance(t *testing.T) {
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

func Test_mapKeyStruct(t *testing.T) {
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

func Test_attrOtherName(t *testing.T) {
	var u1 struct {
		name string "username"
	}
	var u2 struct{ name string }
	var u3 struct{ name string }
	var u4 struct {
		name string "username"
	}

	u1.name = "name1"

	u3.name = "name3"

	//u2 = u1 // Error: cannot use u1 (type struct { name string "username" }) as
	//        type struct { name string } in assignment

	u2 = u3
	fmt.Println(u2.name)

	u4 = u1 //ok
	fmt.Println(u4.name)
}

//
func Test_checkNull(t *testing.T) {
	type User struct {
		id   int
		name string
	}

	user := User{}

	if (user == User{}) {
		fmt.Println("is zero value")
	}
}

func Test_emptyStruct(t *testing.T) {
	a := struct{}{}
	b := struct{}{}

	log.Println(a == b)            // true
	log.Printf("%p, %p\n", &a, &b) // 一樣
}

func Test_pointerStruct(t *testing.T) {
	type ResourcePtr struct {
		id int
	}
	type ResourceVal struct {
		sn int
	}
	type User struct {
		*ResourcePtr
		ResourceVal
		name string
	}
	// u := User{
	// 	&ResourcePtr{1},
	// 	ResourceVal{2},
	// 	"Administrator",
	// }
	u := User{}
	rPtr := ResourcePtr{1}
	rVal := ResourceVal{2}
	u.ResourcePtr = &rPtr
	u.ResourceVal = rVal

	println(u.id)
	println(u.ResourcePtr.id)
	println(u.ResourceVal.sn)

	fmt.Printf("%p %p \n", &rPtr, u.ResourcePtr)
	fmt.Printf("%p %p \n", &rVal, &u.ResourceVal)

	p := uintptr(unsafe.Pointer(&u))
	p += unsafe.Offsetof(u.ResourcePtr.id) //u.id error
	pPtr := (**int)(unsafe.Pointer(p))
	**pPtr = 11

	p = uintptr(unsafe.Pointer(&u))
	p += unsafe.Offsetof(u.sn) //ok
	pVal := (*int)(unsafe.Pointer(p))
	*pVal = 22

	println(u.ResourceVal.sn)
	println(u.id)
	println(u.ResourcePtr.id)
}
