package mathex

import (
	"fmt"
	"testing"
	"math"
)

func Test_abs(t *testing.T) {
	fmt.Println(abs(-3))

}

func Test_max(t *testing.T) {
	max32Add1 := math.MaxInt32 + 1
	max64Add1 := math.MaxInt64 //在這直接+ 1會error
	max64Add1 +=1
	
	fmt.Println(max32Add1)
	fmt.Println(max64Add1)
}
