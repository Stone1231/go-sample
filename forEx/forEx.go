package main

import "fmt"

func main() {
	s := "abc"
	for i, n := 0, len(s); i < n; i++ {
		println(s[i])
	}
	n := len(s)
	for n > 0 { //    while (n > 0) {} //    for (; n > 0;) {}
		n--
		println(s[n])
	}
	// for { //    while (true) {} //    for (;;) {}
	// 	println(s)
	// }

	for i := range s {
		println(s[i]) //  忽略2nd value    支持string/array/slice/map
	}
	for _, c := range s {
		println(c) //    忽略index
	}
	for range s {
		//...
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		println(k, v) // 返回   (key, value)
	}

	a := [3]int{0, 1, 2}
	for i, v := range a { // index value  從複製品取出
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a) //修改成[0, 999, 999]
		}
		a[i] = v + 100 //複製品的value
	}
	fmt.Println(a) //    [100, 101, 102]

	gotoTest()

	breakTest()
}

func gotoTest() {

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
func breakTest() {
L1:
	for x := 0; x < 3; x++ {
	L2:
		for y := 0; y < 5; y++ {
			if y > 2 {
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
