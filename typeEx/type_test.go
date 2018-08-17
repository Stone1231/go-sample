package typeex

import (
	"fmt"
	"math"
	"math/cmplx"
	"testing"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func Test_simple(t *testing.T) {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func Test_int_float(t *testing.T) {
	a1, b1, c1, d1 := 071, 0x1F, 1e9, math.MinInt16
	fmt.Println(a1, b1, c1, d1)
}
