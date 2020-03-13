package testex

import (
	"testing"
)

func ForSlice(s []string) {
	len := len(s)
	for i := 0; i < len; i++ {
		_, _ = i, s[i]
	}
}

func RangeForSlice(s []string) {
	for i, v := range s {
		_, _ = i, v
	}
}

const N = 100000

func initSlice() []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = "www.flysnow.org"
	}
	return s
}

func BenchmarkForSlice(b *testing.B) {
	s := initSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}
func BenchmarkRangeForSlice(b *testing.B) {
	s := initSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSlice(s)
	}
}
