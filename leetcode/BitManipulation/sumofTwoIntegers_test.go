package leetcode

import (
	"fmt"
	"testing"
)

func getSum(a int, b int) int {

	_a := a ^ b
	_b := a & b << 1

	for _a&_b != 0 {
		_a, _b = _a^_b, _a&_b<<1
	}
	_a = _a | _b

	return _a
}

func Test_getSum(t *testing.T) {
	fmt.Println(getSum(2, 3))
	fmt.Println(getSum(18, 19))
	fmt.Println(getSum(20, 30))
	fmt.Println(getSum(-100, 1))	
	fmt.Println(getSum(-7, 3))
}
