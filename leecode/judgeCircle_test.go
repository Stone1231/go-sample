package leecode

import (
	"fmt"
	"testing"
)

func judgeCircle(moves string) bool {

	m := map[byte]int{}

	for index := 0; index < len(moves); index++ {
		m[moves[index]]++
	}

	return m['U'] == m['D'] && m['R'] == m['L']
}

func Test_judgeCircle(t *testing.T) {
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("URRDLL"))
}
