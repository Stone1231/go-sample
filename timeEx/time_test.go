package timeex

import (
	"testing"
	"time"
)

func Test_ex1(t *testing.T) {
	now := time.Now()      // current local time
	sec := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC
	println(sec)
	println(nsec)
}
