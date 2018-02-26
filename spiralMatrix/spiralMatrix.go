package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	// matrix := [][]int{
	// 	[]int{1, 2, 3, 4},
	// 	[]int{5, 6, 7, 8},
	// 	[]int{9, 10, 11, 12},
	// }

	row1 := matrix[0]
	rows := len(matrix)
	cols := len(row1)
	data := make([]int, rows*cols)
	p := &data
	spiralOrder1(0, 0, matrix, p)
	fmt.Println(data[:])
}

func spiralOrder1(idx int, id int, matrix [][]int, data *[]int) {

	if !check(id, matrix) {
		return
	}

	list := matrix[idx]
	count := len(list) - idx

	for index := idx; index < count; index++ {
		(*data)[id] = list[index]
		id++
	}
	//idx++
	spiralOrder2(idx, id, matrix, data)
}

func spiralOrder2(idx int, id int, matrix [][]int, data *[]int) {

	if !check(id, matrix) {
		return
	}

	count := len(matrix)
	col := len(matrix[0]) - (idx + 1)

	for index := idx + 1; index < count; index++ {
		(*data)[id] = matrix[index][col]
		id++
	}
	spiralOrder3(idx, id, matrix, data)
}

func spiralOrder3(idx int, id int, matrix [][]int, data *[]int) {

	if !check(id, matrix) {
		return
	}

	rows := len(matrix)
	list := matrix[rows-(idx+1)]
	count := len(list) - (idx + 1)

	for index := idx; index < count; index++ {
		(*data)[id] = list[count-1-index]
		id++
	}
	//idx++
	spiralOrder4(idx, id, matrix, data)
}

func spiralOrder4(idx int, id int, matrix [][]int, data *[]int) {

	if !check(id, matrix) {
		return
	}

	count := len(matrix) - (idx + 1)

	col := idx
	for index := idx + 1; index < count; index++ {
		(*data)[id] = matrix[count-index][col]
		id++
	}
	idx++
	spiralOrder1(idx, id, matrix, data)
}

func check(id int, matrix [][]int) bool {
	rows := len(matrix)
	row1 := matrix[0]
	cols := len(row1)
	return id < (rows * cols)
}
