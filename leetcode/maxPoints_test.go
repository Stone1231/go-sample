package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

type Point struct {
	X int
	Y int
}

func maxPoints(points []Point) int {

	length := len(points)

	if length == 0 {
		return 0
	} else if length == 1 {
		return 1
	}

	max := 0

	for i := 0; i < length; i++ {
		m := map[string]int{}
		p := 1 //origin point
		max_i := 0
		for j := i + 1; j < length; j++ {
			dx := points[i].X - points[j].X
			dy := points[i].Y - points[j].Y

			if dx == 0 && dy == 0 {
				p++ //duplicate points
				continue
			}

			k := slopeStr(dy, dx)
			m[k]++
			count := m[k]

			if count > max_i {
				max_i = count
			}
		}

		max_i += p

		if max_i > max {
			max = max_i
		}
	}

	return max
}

func slopeStr(dy int, dx int) string {
	if dx == 0 {
		return "1/0"
	} else if dy%dx == 0 {
		return strconv.Itoa(dy / dx)
	} else {
		_dy := int(math.Abs(float64(dy)))
		_dx := int(math.Abs(float64(dx)))

		a := 1
		b := 1
		if _dy > _dx {
			a = _dy
			b = _dx
		} else {
			a = _dx
			b = _dy
		}

		for a%b != 0 {
			_b := b
			b = a % b
			a = _b
		}

		_dy /= b
		_dx /= b

		if (dy < 0 && dx > 0) ||
			(dx < 0 && dy > 0) {
			return "-" + strconv.Itoa(_dy) + "/" + strconv.Itoa(_dx)
		} else {
			return strconv.Itoa(_dy) + "/" + strconv.Itoa(_dx)
		}
	}
}

func Test_maxPoints(t *testing.T) {
	//points := []Point{{1, 1}, {2, 2}, {3, 3}, {4, 2}, {6, 3}}
	//points := []Point{{0, 0}, {0, 0}}
	//points := []Point{{4, 0}, {4, -1}, {4, 5}, {4, -8}}
	//points := []Point{{1, 1}, {1, 1}, {1, 1}}
	//points := []Point{{0, 0}, {94911151, 94911150}, {94911152, 94911151}}
	points := []Point{{3, 1}, {12, 3}, {3, 1}, {-6, -1}}
	fmt.Println(maxPoints(points))

	// fmt.Println(slope_str(1, 2))
	// fmt.Println(slope_str(3, 6))
	// fmt.Println(slope_str(6, 3))
	// fmt.Println(slope_str(-4, 2))
	// fmt.Println(slope_str(4, -2))
	//fmt.Println(slope_str(2, -4))
	//fmt.Println(slope_str(-2, 4))
	//fmt.Println(slope_str(10, 5))
}
