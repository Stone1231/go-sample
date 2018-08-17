package errorex

import (
	"fmt"
	"testing"
)

func Test_error(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			print("recover:")
			println(err.(string)) //interface{} 轉型為具體類型
		}
	}()
	panic("panic error!")
}

///只有最後的錯誤可被捕獲
func Test_error2(t *testing.T) {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

func Test_error3(t *testing.T) {
	x, y := 6, 4
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

func Test_ErrDivByZero(t *testing.T) {
	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}

func Test_callError(t *testing.T) {

	for _, i := range []int{7, 42} {
		if r, e := f42Error(i); e != nil {
			fmt.Println("f42Error failed:", e)
		} else {
			fmt.Println("f42Error worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f42CustomError(i); e != nil {
			fmt.Println("f42CustomError failed:", e)
		} else {
			fmt.Println("f42CustomError worked:", r)
		}
	}

	// If you want to programmatically use the data in
	// a custom error, you'll need to get the error  as an
	// instance of the custom error type via type
	// assertion.
	_, e := f42CustomError(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
