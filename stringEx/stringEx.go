package stringEx

import (
	"bytes"
)

func reverseStr(s string) string {
	var buffer bytes.Buffer
	l := len(s)
	for index := l - 1; index >= 0; index-- {
		buffer.WriteByte(s[index])
	}
	return buffer.String()
}
