package dsa

import (
	"fmt"
	"testing"
)

func Test_array1(t *testing.T) {
	a := []int{0, 0, 0}
	a[1] = 10
	b := make([]int, 3) //makeslice
	b[1] = 10
	//c := new([]int)
	//c[1] = 10  Error: invalid operation: c[1] (index of type *[]int)
	fmt.Println(a[1], b[1])
}

//array是值類型
func Test_array2(t *testing.T) {
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

func Test_array2DInit(t *testing.T) {
	n := 3
	m := 2
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}
	fmt.Println(res)
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

func Test_randomArray(t *testing.T) {
	fmt.Println(GetRandArray(100, 8))
}

func Test_randomArray2D(t *testing.T) {
	fmt.Println(GetRandArray2D(50, 100, 8))
}
