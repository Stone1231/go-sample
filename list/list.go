package main

import "fmt"

func main() {
	// array1()

	// array2()

	// array3()

	// array4()

	// array5()

	//slice1()

	// slice2()

	//slice3()

	//slice4()

	//slice5()

	//reslice1()

	//append1()

	// append2()
	// append3()

	//append4()

	//copy1()

	//map1()

	//map2()

	//map3()

	map4()
}

func array1() {
	a := []int{0, 0, 0}
	a[1] = 10
	b := make([]int, 3) //makeslice
	b[1] = 10
	//c := new([]int)
	//c[1] = 10  Error: invalid operation: c[1] (index of type *[]int)
	fmt.Println(a[1], b[1])
}

///array是值類型
func array2() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 200}
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(a[2], b[2], c[2], d[1].name)
}

///多維
func array3() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} //   第2維度不能用"..."
	fmt.Println(a[1][1], b[2][1])
}

///copy會影響效能
func copyArray4(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}
func array4() {
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	copyArray4(a)
	fmt.Println(a)
}

///count
func array5() {
	a := [2]int{}
	println(len(a), cap(a)) // 2, 2
}

//slice
func slice1() {

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(data[1:4:5]) // [low : high : max]
	s := data[1:4:5]
	fmt.Println(s, len(s), cap(s))

	for index := 0; index < len(s); index++ {
		fmt.Println(index, s[index])
	}
	// for var v:= range s {
	// 	fmt.Println(v)
	// }

	fmt.Println(data[:6:8])
	fmt.Println(data[5:])
	fmt.Println(data[:3])
	fmt.Println(data[:])
}

func slice2() {
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	fmt.Println(data)
}

func slice3() {
	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))
}

func slice4() {
	s := []int{0, 1, 2, 3}

	p := &s[2] // *int
	*p += 100
	fmt.Println(s)
}

func slice5() {
	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data[0][0])

	d := [5]struct {
		x int
	}{}
	s := d[:]
	d[1].x = 10
	s[2].x = 20
	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
}

///reslice
func reslice1() {
	// s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s1 := s[2:5]    //[2 3 4]
	// s2 := s1[2:6:7] //[4 5 6 7]
	// s3 := s2[3:6] //error

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5] // [2 3 4]
	s1[2] = 100
	s2 := s1[2:6] // [100 5 6 7]
	s2[3] = 200
	fmt.Println(s)
}

//append
func append1() {
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)
	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s, s2)
}

func append2() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := data[:3]
	s2 := append(s, 100, 200) //
	fmt.Println(data)
	fmt.Println(s)
	fmt.Println(s2)

	fmt.Println(&s[0], &s2[0], &data[0]) //一樣
}

func append3() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	s = append(s, 100, 200) //超過s.cap 重新分配
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0]) //不一樣
}

func append4() { //通常以兩倍容量重新分配底層數組
	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}

///copy
func copy1() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:] //index8個以後
	fmt.Println(s)
	s2 := data[:5] //index不超過5
	fmt.Println(s2)
	copy(s2, s) // dst:s2, src:s   from s to s2
	fmt.Println(s2)
	fmt.Println(data)
}

///map
func map1() {
	m := map[int]struct {
		name string
		age  int
	}{
		1: {"user1", 10},
		2: {"user2", 20},
	}
	println(m[1].name)
}

func map2() {

	//m := make(map[string]int, 1000)

	m := map[string]int{
		"a": 1,
	}

	if v, ok := m["a"]; ok { //contains a key in go?
		println(ok) //bool
		println(v)
	}
	println(m["c"])
	m["b"] = 2
	delete(m, "c")
	println(len(m))
	for k, v := range m {
		println(k, v)
	}
}

func map3() {
	type user struct{ name string }
	m := map[int]user{
		1: {"user1"},
	}
	//m[1].name = "Tom" // Error: cannot assign to m[1].name, map not addressable

	u := m[1]
	u.name = "Tom"
	m[1] = u //    value
	fmt.Println(m[1].name)

	m2 := map[int]*user{
		1: &user{"user1"},
	}
	m2[1].name = "Jack" //指標副本 可以透過指標修改
}

func map4() { //泆代時可刪除鍵值 但期間有新增操作會有意外
	for i := 0; i < 5; i++ {
		m := map[int]string{
			0: "a", 1: "a", 2: "a", 3: "a", 4: "a",
			5: "a", 6: "a", 7: "a", 8: "a", 9: "a",
		}
		for k := range m {
			m[k+k] = "x"
			delete(m, k)
		}
		fmt.Println(m)
	}
}
