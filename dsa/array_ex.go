package dsa

import "fmt"

func array_ex1() {
	a := []int{0, 0, 0}
	a[1] = 10
	b := make([]int, 3) //makeslice
	b[1] = 10
	//c := new([]int)
	//c[1] = 10  Error: invalid operation: c[1] (index of type *[]int)
	fmt.Println(a[1], b[1])
}

///array是值類型
func array_ex2() {
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
func array_ex3() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} //   第2維度不能用"..."
	fmt.Println(a[1][1], b[2][1])
}

///copy會影響效能
func array_copy(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}
func array_ex4() {
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	array_copy(a)
	fmt.Println(a)
}

///count
func array_ex5() {
	a := [2]int{}
	println(len(a), cap(a)) // 2, 2
}
