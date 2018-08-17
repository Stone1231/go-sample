package dsa

import "fmt"

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
