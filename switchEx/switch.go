package main

func main() {
	x := []int{1, 2, 3}
	i := 2
	switch i {
	case x[1]:
		println("a")
	case 1, 3:
		println("b")
	default:
	}
	println("c")

	x1 := 10
	switch x1 {
	case 10:
		println("a1")
		fallthrough //強制跳到下一個
	case 0:
		println("b1")
	case 20:
		println("c1")
	}

	switch {
	case x[1] > 0:
		println("a2")
	case x[1] < 0:
		println("b2")
	default:
		println("c2")
	}
	switch i := x[2]; {//帶初始化語法
	case i > 0:
		println("a3")
	case i < 0:
		println("b3")
	default:
		println("c3")
	}
}
