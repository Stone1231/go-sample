package converEx

import (
	"fmt"
	"testing"
)

func Test_ByteToInt(t *testing.T) {
	var b byte = 100
	// var n int = b // Error: cannot use b (type byte) as type int in assignment
	var n int = int(b) //
	fmt.Println(n)
}
