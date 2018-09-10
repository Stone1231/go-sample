package leetcode

import (
	"fmt"
	"testing"
)

func generateMatrix(n int) [][]int {

	left, right, top, bottom := 0, n-1, 0, n-1

	i := 0

	res := make([][]int, n)
	for index := range res {
		res[index] = make([]int, n)
	}

	s := 0
	for left <= right {
		s = left
		for s <= right {
			i++
			res[top][s] = i
			s++
		}
		top++

		s = top
		for s <= bottom {
			i++
			res[s][right] = i
			s++
		}
		right--

		s = right
		for s >= left {
			i++
			res[bottom][s] = i
			s--
		}
		bottom--

		s = bottom
		for s >= top {
			i++
			res[s][left] = i
			s--
		}
		left++
	}
	return res
}

func spiralOrder(matrix [][]int) []int {

	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}
	row1 := matrix[0]
	cols := len(row1)

	left, right, top, bottom := 0, cols-1, 0, rows-1

	res := []int{}

	s := 0
	for {
		s = left
		for s <= right {
			res = append(res, matrix[top][s])
			s++
		}
		top++
		if top > bottom {
			break
		}

		s = top
		for s <= bottom {
			res = append(res, matrix[s][right])
			s++
		}
		right--
		if left > right {
			break
		}

		s = right
		for s >= left {
			res = append(res, matrix[bottom][s])
			s--
		}
		bottom--
		if top > bottom {
			break
		}

		s = bottom
		for s >= top {
			res = append(res, matrix[s][left])
			s--
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

func Test_matrix(t *testing.T) {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
}

func Test_spiralOrder(t *testing.T) {
	fmt.Println(spiralOrder(
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		}))

	fmt.Println(spiralOrder(
		[][]int{
			[]int{2, 3, 4},
			[]int{5, 6, 7},
			[]int{8, 9, 10},
			[]int{11, 12, 13},
		}))

	fmt.Println(spiralOrder(
		[][]int{
			[]int{6, 9, 7},
		}))

	fmt.Println(spiralOrder(
		[][]int{
			[]int{2},
			[]int{5},
			[]int{8},
			[]int{11},
		}))
}
