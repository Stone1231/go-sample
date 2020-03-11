package dsa

import (
	"fmt"
	"testing"
)

//array是值類型
func Test_simple(t *testing.T) {
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

//多維
func Test_array2d(t *testing.T) {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} //   第2維度不能用"..."
	fmt.Println(a[1][1], b[2][1])
}
func Test_arrayValPtr(t *testing.T) {
	a := [4]int{}
	fmt.Printf("a: %p\n", &a)
	arrayVal(a, 1, 1000)
	arrayPtr(&a, 1, 2000)
	fmt.Println(a)
}
func Test_ArrayLen(t *testing.T) {
	a := [2]int{}
	println(len(a), cap(a)) // 2, 2
}
func Test_arrayPtr(t *testing.T) {
	a := [4]int{}
	//a := make([]int, 4)
	a[0] = 2
	fmt.Printf("a: %p\n", &a[0])
	b := a
	fmt.Printf("b: %p\n", &b[0])
	c := &a
	fmt.Printf("c: %p\n", c)
	fmt.Printf("%v \n", c[0])
	for i := range c {
		fmt.Printf("%v \n", c[i])
	}
}
