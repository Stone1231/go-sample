package reflectex

import (
	"fmt"
	"reflect"
	"testing"
)

//不同類型的值不會深度相等
func Test_DiffStruct(t *testing.T) {
	type S1 struct {
		Field int
	}
	type S2 struct {
		Field int
	}
	fmt.Println(reflect.DeepEqual(S1{1}, S2{1}))
}

func Test_array(t *testing.T) {
	a1 := [...]string{"hello1", "hello2"}
	a2 := [...]string{"hello1", "hello2"}
	fmt.Println(reflect.DeepEqual(a1, a2))
}

func Test_struct(t *testing.T) {
	type S struct {
		Field1 int
		field2 string
	}
	s1 := S{Field1: 1, field2: "hello"}
	s2 := S{Field1: 1, field2: "hello"}
	fmt.Println(reflect.DeepEqual(s1, s2))
}

func Test_func(t *testing.T) {
	f1 := func(a int) int {
		return a * 2
	}
	fmt.Println(reflect.DeepEqual(f1, f1))
	f1 = nil
	fmt.Println(reflect.DeepEqual(f1, f1))
}

func Test_interface(t *testing.T) {
	var i1 interface{}
	i1 = "hello"
	var i2 interface{}
	i2 = "hello"
	fmt.Println(reflect.DeepEqual(i1, i2))
}

// 1.兩個map都為nil或者都不為nil，並且長度要相等
// they are both nil or both non-nil, they have the same length
// 2.相同的map對像或者所有key要對應相同
// either they are the same map object or their corresponding keys
// 3.map對應的value也要深度相等
func Test_map(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(reflect.DeepEqual(m1, m2))
}

// 1.兩個指針滿足go的==操作符
//   Pointer values are deeply equal if they are equal using Go's == operator
// 2.兩個指針指向的值是深度相等的
func Test_ptr(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	M1 := &m1
	M2 := &m2
	fmt.Println(reflect.DeepEqual(M1, M2))
}

// 1.兩個切片都為nil或者都不為nil，並且長度要相等
// they are both nil or both non-nil, they have the same length
// 2.兩個切片底層數據指向的第一個位置要相同或者底層的元素要深度相等
// either they point to the same initial entry of the same underlying array (that is, &x[0] == &y[0]) or their corresponding elements (up to length) are deeply equal.
func Test_slice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[0:3]
	s3 := s1[0:3]
	fmt.Println(reflect.DeepEqual(s2, s3))
	s4 := s1[1:4]
	fmt.Println(reflect.DeepEqual(s2, s4))
}

//空的切片跟nil切片是不深度相等的，例如
func Test_nil(t *testing.T) {
	s1 := []byte{}
	s2 := []byte(nil)
	fmt.Println(reflect.DeepEqual(s1, s2))
}
