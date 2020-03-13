package testex

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
	}{
		{
			name: "1+1",
			a:    1,
			b:    1,
			c:    2,
		}, {
			name: "2+3",
			a:    2,
			b:    3,
			c:    4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.c == (tt.a + tt.b) {
				t.Logf("%v + %v = %v", tt.a, tt.b, tt.c)
			} else {
				t.Errorf("%v + %v != %v", tt.a, tt.b, tt.c)
			}
		})
	}
}
