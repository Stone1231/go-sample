package funcex

import (
	"fmt"
	"os"
)

func sprintfEx(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func inputFn(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func inputParams(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

func twoNums() (int, int) {
	return 1, 2
}

func add(x, y int) int {
	return x + y
}

func addReturnZ(x, y int) (z int) {
	z = x + y
	return
}

func add100Defer(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	return
}

func add200DeferPrint(x, y int) (z int) {
	defer func() {
		println(z)
	}()
	z = x + y
	return z + 200
	//順序: (z = z + 200) -> (call defer) -> (return)
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}

///延遲引用
func delay() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

///
func fileWriteDeferClose() error {
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString("Hello, World2!")
	return nil
}
