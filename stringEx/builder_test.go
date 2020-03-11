package stringEx

import (
	"fmt"
	"strings"
	"testing"
)

func Test_builder(t *testing.T) {
	var b strings.Builder
	write := func(msg string) {
		_, _ = b.WriteString(msg)
		_, _ = b.WriteRune('\n')
	}
	write("a")
	write("b")
	write("c")
	fmt.Println(b.String())
}
