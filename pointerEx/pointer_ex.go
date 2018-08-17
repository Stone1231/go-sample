package pointerex

import "fmt"

func zeroptr(iptr *int) {
	fmt.Println("*iptr:", *iptr) // = i

	*iptr = 0

	fmt.Println("iptr:", iptr) // = &i
}

func zeroval(ival int) {
	ival = 0

	fmt.Println("&ival:", &ival)
}
