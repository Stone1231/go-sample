package loopex

import (
	"fmt"
	"testing"
)

func Test_str(t *testing.T) {
	s := "abc"

	for i, n := 0, len(s); i < n; i++ {
		println(s[i])
	}

	n := len(s)
	for n > 0 { //    while (n > 0) {} //    for (; n > 0;) {}
		n--
		println(s[n])
	}

	for i := range s {
		println(s[i]) //  忽略2nd value    支持string/array/slice/map
	}

	for _, c := range s {
		println(c) //    忽略index
	}

	for range s {
		//...
	}
}
func Test_map(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		println(k, v) // 返回   (key, value)
	}
}
func Test_array(t *testing.T) {

	a := [3]int{0, 1, 2}
	for i, v := range a { // index value  從複製品取出
		if i == 0 {
			a[1], a[2] = 99, 999
			fmt.Println(a) //[0, 99, 999]
		}
		a[i] = v + 100 //複製品的value
	}
	fmt.Println(a) //[100, 101, 102]
}
func Test_slice(t *testing.T) {

	s := []int{0, 1, 2}
	for i, v := range s {
		if i == 0 {
			s[1], s[2] = 99, 999
			fmt.Println(s) //[0, 99, 999]
		}
		s[i] = v + 100 //會隨著s變化的value
	}
	fmt.Println(s) //[100 199 1099]
}

func Test_goto(t *testing.T) {
	var i int
	for {
		println(i)
		i++
		if i > 2 {
			goto BREAK
		}
	}
	goto EXIT

BREAK:
	println("break")
EXIT:
	println("exit")
	// Error: label EXIT defined and not used

}

//break可用在for switch select, 而continue只能用在for迴圈
func Test_break(t *testing.T) {
L1:
	for x := 0; x < 3; x++ {
	L2:
		for y := 0; y < 5; y++ {
			if y == 2 {
				continue L2
			}
			if x > 1 {
				break L1
			}
			print(x, ":", y, " ")
		}
		println()
	}
}
