package dsa

import (
	"fmt"
	"math/rand"
)

//copy會影響效能
func arrayVal(a [4]int, i int, v int) {
	//fmt.Printf("%p\n", &a)
	a[i] = v
}

func arrayPtr(a *[4]int, i int, v int) {
	//fmt.Printf("ptr: %p\n", a)
	a[i] = v
}

func printArrayItemAddress(a *[4]int) {
	for index := 0; index < len(*a); index++ {
		fmt.Printf("&array-%v %p \n", index, &(*a)[index])
	}
	fmt.Println()
}

func GetRandArray(length int, max int) []int {
	a := make([]int, length)

	for index := 0; index < length; index++ {
		a[index] = rand.Intn(max)
	}

	return a
}

func GetRandArray2D(n int, m int, max int) [][]int {
	aa := make([][]int, n)

	for i := 0; i < n; i++ {
		a := GetRandArray(m, max)
		aa[i] = a
	}

	return aa
}
