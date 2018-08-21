package leecode

import (
	"fmt"
	"testing"
)

type Matrix struct {
	array  *[][]int
	idx    int
	rows   int
	cols   int
	length int
	data   []int
	id     int
}

func (self *Matrix) insert(item int) {
	self.data[self.id] = item
	self.id++
}

func spiralOrderOld(matrix [][]int) []int {

	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}
	row1 := matrix[0]
	cols := len(row1)
	length := rows * cols
	m := Matrix{
		array:  &matrix,
		idx:    0,
		rows:   rows,
		cols:   cols,
		length: length,
		data:   make([]int, length),
		id:     0,
	}

	top(&m)
	return m.data
}

//func top(idx int, id int, matrix [][]int, data *[]int) {
func top(m *Matrix) {

	if !check(m) {
		return
	}

	list := (*m.array)[m.idx] //matrix[idx]
	count := m.cols - m.idx*2 //len(list) - idx*2
	add := m.idx

	for index := 0; index < count; index++ {
		m.insert(list[index+add])
	}

	right(m)
}

func right(m *Matrix) {

	if !check(m) {
		return
	}

	count := m.rows - (2*m.idx + 1)
	col := m.cols - (m.idx + 1)
	add := m.idx + 1

	for index := 0; index < count; index++ {
		m.insert((*m.array)[index+add][col])
	}

	bottom(m)
}

func bottom(m *Matrix) {

	if !check(m) {
		return
	}

	//rows := len(matrix)
	list := (*m.array)[m.rows-(m.idx+1)]
	length := m.cols
	count := length - (2*m.idx + 1)
	add := m.idx + 1

	for index := 0; index < count; index++ {
		m.insert(list[length-1-index-add])
	}

	left(m)
}

func left(m *Matrix) {

	if !check(m) {
		return
	}

	length := m.rows //len(matrix)
	count := length - (2 * (m.idx + 1))
	col := m.idx
	add := m.idx + 1

	for index := 0; index < count; index++ {
		m.insert((*m.array)[length-1-index-add][col])
	}

	m.idx++
	top(m)
}

func check(m *Matrix) bool {
	return m.id < (m.length)
}

func Test_spiralOrderOld(t *testing.T) {
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

	// error
	// fmt.Println(spiralOrder(
	// 	[][]int{
	// 		[]int{6, 9, 7},
	// 	}))
}
