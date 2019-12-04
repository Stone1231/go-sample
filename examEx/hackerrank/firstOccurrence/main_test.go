package main

import (
	"fmt"
	"testing"
)

func firstOccurrence(s string, x string) int32 {

	count := len(x)
	sl := len(s)
	start := 0
	for i := 0; i < count; i++ {
		index := start + i
		if index >= sl {
			start = -1
			break
		}
		if x[i] != '*' && s[index] != x[i] {
			start += (i + 1)
			i = -1
		}
	}

	return int32(start)
}

func Test_exam(t *testing.T) {
	fmt.Println(firstOccurrence("juliasamanthantjulia", "ant"))
}
