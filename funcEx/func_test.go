package funcex

import (
	"fmt"
	"testing"
	"time"
)

func Test_sprintfEx(t *testing.T) {
	n, s := sprintfEx(1, 2, "%T") //%v
	fmt.Println(n, s)
}

func Test_inputFn(t *testing.T) {
	n := inputFn(func() int { return 100 })
	ss := format(
		func(s string, x, y int) string {
			return fmt.Sprintf(s, x, y)
		},
		"%d, %d",
		10,
		20)
	println(n, ss)
}

func Test_inputParams(t *testing.T) {
	println(inputParams("sum: %d", 1, 2, 3))

	slice1 := []int{1, 2, 3} //不能用array
	println(inputParams("sum: %d", slice1...))
}

func Test_addSum_defer(t *testing.T) {
	//x, y := twoNums() //y declared and not used
	x, _ := twoNums()
	println(x)
	println(add(twoNums()))
	println(sum(twoNums()))
	println(addReturnZ(2, 3))
	println(add100Defer(3, 4))
	add200DeferPrint(4, 5)
}

func Test_functionValue(t *testing.T) {
	fn := func() { println("Hello, World!") }
	fn()

	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](100))

	d := struct {
		fn1 func() string
		fn2 func() string
	}{
		fn1: func() string { return "Hello, fn1!" },
		fn2: func() string { return "Hello, fn2!" },
	}
	println(d.fn1())

	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, fc!" }
	println((<-fc)())
}

func Test_delay(t *testing.T) {
	delay()

	time.Sleep(1 * time.Second)

	println("use delay:")
	f := delay()
	f()
}

func Test_fileWrite(t *testing.T) {
	res := fileWriteDeferClose()

	if res != nil {
		println("error")
	}
}

func Test_defer(t *testing.T) {
	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, y) // y封包引用
	}(x) //雖然defer最後執行但 x已先複製
	x += 10
	y += 100
	println("x =", x, "y =", y)
}
