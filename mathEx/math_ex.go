package mathex

import "math"

func abs(n int64) int64 {
	return int64(math.Abs(float64(n)))
}
