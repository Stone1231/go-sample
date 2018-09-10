package leetcode

import (
	"fmt"
	"github.com/sample/dsa"
	"testing"
	"time"
)

func setZeroes(matrix [][]int) {

	rows := len(matrix)
	if rows == 0 {
		return
	}
	row1 := matrix[0]
	cols := len(row1)

	zeroCols := map[int]int{}
	zeroRow := false

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				if !zeroRow {
					for k := 0; k < j; k++ {
						matrix[i][k] = 0
					}
					zeroRow = true
				}
				zeroCols[j] = 1
			} else if zeroRow {
				matrix[i][j] = 0
			}
		}
		zeroRow = false
	}

	for k := range zeroCols {
		for i := 0; i < rows; i++ {
			matrix[i][k] = 0
		}
	}
	//fmt.Println(matrix)
}

func setZeroesHis(matrix [][]int) {

	rows := len(matrix)
	if rows == 0 {
		return
	}
	row1 := matrix[0]
	cols := len(row1)

	zeroCols := map[int]int{}
	zeroRow := false

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				zeroRow = true
				zeroCols[j]++
			}
		}
		if zeroRow {
			for j := 0; j < cols; j++ {
				matrix[i][j] = 0
			}
		}
		zeroRow = false
	}

	for i := 0; i < rows; i++ {
		for k := range zeroCols {
			matrix[i][k] = 0
		}
	}

	//fmt.Println(matrix)
}

func Test_setZeroes(t *testing.T) {
	// setZeroes(
	// 	[][]int{
	// 		[]int{1, 1, 1},
	// 		[]int{1, 0, 1},
	// 		[]int{1, 1, 1},
	// 	})

	aa := dsa.GetRandSlice2D(5000, 5000, 10)
	t1 := time.Now()
	setZeroes(aa)
	elapsed := time.Since(t1)
	fmt.Println("time1: ", elapsed)

	aa = dsa.GetRandSlice2D(5000, 5000, 10)
	t1 = time.Now()
	setZeroesHis(aa)
	elapsed = time.Since(t1)
	fmt.Println("time2: ", elapsed)
}
