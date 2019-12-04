package converex

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_ByteToInt(t *testing.T) {
	var b byte = 100
	// var n int = b // Error: cannot use b (type byte) as type int in assignment
	var n int = int(b) //
	fmt.Println(n)
}

func Test_strToInt(t *testing.T) {
	str := "1234"

	/** converting the str1 variable into an int using Atoi method */
	i1, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println(i1)
	}
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
