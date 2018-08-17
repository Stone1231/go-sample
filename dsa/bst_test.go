package dsa

import (
	"fmt"
	"testing"
)

///sort 二元樹排序
func Test_sort(t *testing.T) {
	a := []int{3, 2, 1, 7, 8, 4}
	Sort(a)
	fmt.Println(a)
}

func Test_sortAppend(t *testing.T) {
	var root *tree
	root = add(root, 3)
	root = add(root, 2)
	root = add(root, 1)
	root = add(root, 7)

	a := []int{} // = make([]int, 0)
	appendValues(&a, root)
	fmt.Println(a)

	b := []int{} //make([]int, 0, 4)
	b2 := appendValuesOld(b, root)
	fmt.Println(b2)

	c := make([]int, 4)
	//c[:0]算另一個slice會影響到c
	appendValuesOld(c[:0], root)
	fmt.Println(c)
}
