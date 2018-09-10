package dsa

import "fmt"
import "math/rand"

func sliceVal(s []int, i int, v int) {
	//fmt.Printf("%p\n", &s)
	s[i] = v
}

func printSliceItemAddress(s *[]int) {
	for index := 0; index < len(*s); index++ {
		fmt.Printf("&slice-%v %p \n", index, &(*s)[index])
	}
	fmt.Println()
}

func GetRandSlice(length int, max int) []int {
	s := make([]int, length)

	for index := 0; index < length; index++ {
		s[index] = rand.Intn(max)
	}

	return s
}

func GetRandSlice2D(n int, m int, max int) [][]int {
	ss := make([][]int, n)

	for i := 0; i < n; i++ {
		s := GetRandSlice(m, max)
		ss[i] = s
	}

	return ss
}

func CopySlice2D(ss [][]int) [][]int{

	_ss := make([][]int, len(ss))			
	
	for i, s := range ss {

		//_ss[i] = append(_ss[i], s...)

		//同上, 效能較好
		_ss[i] = make([]int, len(s))
		copy(_ss[i], s)
	}

	return _ss
}
