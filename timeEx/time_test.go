package timeex

import (
	"fmt"
	"testing"
	"time"
)

func Test_simple(t *testing.T) {
	now := time.Now()      // current local time
	sec := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC
	println(sec)
	println(nsec)
}

func Test_timeFormat(t *testing.T) {
	var formats = []string{
		"2006-1-2T15:4:5.999999999Z07:00", // RCF3339Nano with short date fields.
		"2006-1-2t15:4:5.999999999Z07:00", // RFC3339Nano with short date fields and lower-case "t".
		"2006-1-2 15:4:5.999999999",       // space separated with no time zone
		"2006-1-2",                        // date only
		// Notable exception: time.Parse cannot handle: "2001-12-14 21:59:43.10 -5"
		// from the set of examples.
	}

	var ss = []string{
		"1989-1-2T15:4:5.999999999Z07:00", // RCF3339Nano with short date fields.
		"1989-1-2t15:4:5.999999999Z07:00", // RFC3339Nano with short date fields and lower-case "t".
		"1989-1-2 15:4:5.999999999",       // space separated with no time zone
		"1989-1-2",                        // date only
	}

	for i := 0; i < len(formats); i++ {
		if t, err := time.Parse(formats[i], ss[i]); err == nil {
			fmt.Printf("%v\n", t)
		} else {
			fmt.Println(err.Error())
		}
	}

	now := time.Now()
	for i := 0; i < len(formats); i++ {
		timeFormat := now.Format(formats[i])
		t, err := time.Parse(formats[i], timeFormat)
		if t.Year() == 0 || err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(timeFormat)
		}
	}
}
