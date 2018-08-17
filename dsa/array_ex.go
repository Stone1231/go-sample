package dsa

import "fmt"

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
