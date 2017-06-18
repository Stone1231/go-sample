package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// n, s := test1(1, 2, "%T") //%v
	// fmt.Println(n, s)

	// s1 := test2(func() int { return 100 })
	// s2 := format(func(s string, x, y int) string {
	// 	return fmt.Sprintf(s, x, y)
	// }, "%d, %d", 10, 20)
	// println(s1, s2)

	// println(test3("sum: %d", 1, 2, 3))
	// s3 := []int{1, 2, 3}
	// println(test3("sum: %d", s3...))

	// x, _ := test4()
	// println(x)
	// println(add(test4()))
	// println(sum(test4()))

	// println(add2(2, 3))

	// println(add3(3, 4))

	// add4(4, 5)

	// funVal()

	// f := test5()
	// f()

	// defer1()

	// defer2(0) // 100/0 error

	// defer3()

	// errorEx()

	// errorEx2()

	//errorEx3()

	//errorEx4(6, 4)

	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}

func test1(x, y int, s string) (int, string) { //
	n := x + y //
	return n, fmt.Sprintf(s, n)
}

///
func test2(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string //

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

///
func test3(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

///
func test4() (int, int) {
	return 1, 2
}

func add(x, y int) int {
	return x + y
}
func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}

///
func add2(x, y int) (z int) {
	z = x + y
	return
}

///
func add3(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	return
}

///
func add4(x, y int) (z int) {
	defer func() {
		println(z)
	}()
	z = x + y
	return z + 200 //順序: (z = z + 200) -> (call defer) -> (ret)
}

///
func funVal() {
	fn := func() { println("Hello, World1!") }
	fn()

	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](100))

	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World2!" },
	}
	println(d.fn())

	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World3!" }
	println((<-fc)())
}

///延遲引用
func test5() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

///
func defer1() error {
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString("Hello, World2!")
	return nil
}

///
func defer2(x int) {
	defer println("a")
	defer println("b")
	defer func() {
		//println(100 / x)
	}()
	defer println("c")
}

///
func defer3() {
	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, y) // y封包引用
	}(x) // x複製
	x += 10
	y += 100
	println("x =", x, "y =", y)
}

///
func errorEx() {

	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) //   interface{} 轉型為具體類型
		}
	}()
	panic("panic error!")
}

///只有最後的錯誤可被捕獲
func errorEx2() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

///
func errorEx3() {
	defer recover() //無效
	//defer fmt.Println(recover().(string)) //無效
	defer func() {
		func() {
			//println("defer inner")
			//println(recover().(string)) //無效
		}()
	}()

	defer except()

	panic("test panic")
}
func except() {
	println(recover().(string))
}

///
func errorEx4(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
		return
	}()
	println("x / y =", z)
}

///內部使用panic  對外API使用error
var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}
