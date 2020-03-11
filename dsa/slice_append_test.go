package dsa

import (
	"fmt"
	"testing"
)

func Test_appendSimple(t *testing.T) {
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)
	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s, s2)
}

func Test_appendIndexAddress(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := data[:3]
	s2 := append(s, 100, 200) //
	fmt.Println(data)
	fmt.Println(s)
	fmt.Println(s2)

	fmt.Println(&s[0], &s2[0], &data[0]) //一樣
}

func Test_appendCapOverflow(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	s = append(s, 200)
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0]) //一樣

	s = append(s, 300, 400) //超過s.cap 重新分配
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0]) //不一樣
}

//通常以兩倍容量重新分配底層數組
func Test_appendCap(t *testing.T) {
	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d , %p\n", c, n, &s)
			c = n
		}
	}
}

func Test_insertFirst(t *testing.T) {
	s := []int{2, 3, 4, 5}
	s = append([]int{1}, s...)
	fmt.Println(s)
}

func Test_insertAtIndex(t *testing.T) {
	s := []int{1, 2, 4, 5}
	i := 2
	val := 3
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = val

	fmt.Println(s)
}

func Test_appendSlice(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := append(s1[:0:0], s1...)
	s3 := append([]int{}, s1...)

	s1[0] = 10
	s2[2] = -2
	s3[3] = -3

	//都不會互相影響
	fmt.Printf("s1: %v \n", s1)
	fmt.Printf("s2: %v \n", s2)
	fmt.Printf("s3: %v \n", s3)
}

func Test_appendSlice2(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{5, 6, 7, 8, 9}
	s1 = append(s1, s2...)

	s1[5] = 60
	s2[2] = 70

	//都不會互相影響
	fmt.Printf("s1: %v \n", s1)
	fmt.Printf("s2: %v \n", s2)
}

func Test_appendSlice2D(t *testing.T) {
	ss1 := [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}}
	//2D會互相影響
	//ss2 := append(ss1[:0:0], ss1...)

	//特別處理才不會互相干擾
	ss2 := CopySlice2D(ss1)

	ss1[0][0] = 10
	ss2[1][1] = 20

	fmt.Printf("s1: %v \n", ss1)
	fmt.Printf("s2: %v \n", ss2)
}
