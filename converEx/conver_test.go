package converex

import (
	"fmt"
	"testing"
	"strconv"
)

func Test_ByteToInt(t *testing.T) {
	var b byte = 100
	// var n int = b // Error: cannot use b (type byte) as type int in assignment
	var n int = int(b) //
	fmt.Println(n)
}

func Test_intToStr(t *testing.T) {
	n := 11
	var s string = "it is " + strconv.Itoa(n)
	fmt.Println(s)
}

func Test_arrayToString(t *testing.T) {
    a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println(ArrayToString(a, ",")) //1,2,3,4,5,6,7,8,9
}