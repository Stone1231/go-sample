package switchex

import "testing"

var x = []int{1, 2, 3}

func Test_simple(t *testing.T) {

	i := 2
	switch i {
	case x[1]:
		println("x[1]")
	case 1, 3:
		println("1, 3")
	default:
	}
	println("end")
}
func Test_fallthrough(t *testing.T) {
	v := 10
	switch v {
	case 10:
		println("10")
		fallthrough //強制跳到下一個
	case 0:
		println("next")
	case 20:
		println("20")
	}
}
func Test_condition(t *testing.T) {
	switch {
	case x[1] > 0:
		println("> 0")
	case x[1] < 0:
		println("< 0")
	default:
		println("default")
	}
}

func Test_init(t *testing.T) {
	switch i := x[2]; { //帶初始化語法
	case i > 0:
		println("i > 0")
	case i < 0:
		println("i < 0")
	default:
		println("c3")
	}
}
