package dsa

import (
	"fmt"
	"testing"
)

func Test_sliceRange(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[1:4:5] // [low : high : max]

	fmt.Println(s, len(s), cap(s)) //[1 2 3] 3 4

	fmt.Println(data[6:8]) //[6 7]
	fmt.Println(data[5:])  //[5 6 7 8 9]
	fmt.Println(data[:5])  //[0 1 2 3 4]
	fmt.Println(data[:0])  //[]
	fmt.Println(data[:])   //[0 1 2 3 4 5 6 7 8 9]
}

func Test_updateSliceToArray(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	fmt.Println(data)
}

func Test_sliceInit(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))
	// 0 1 2 3 4 5 6 7 8
	//[0 1 2 3 0 0 0 0 100] 9 9

	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))
	//[0 0 0 0 0 0] 6 8

	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))
	//[0 0 0 0 0 0] 6 6
}

func Test_sliceItemPtr(t *testing.T) {
	s := []int{0, 1, 2, 3}

	p := &s[2] // *int
	*p += 100
	fmt.Println(s)
}

func Test_slice2D(t *testing.T) {
	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data[0][0])

	d := [5]struct {
		x int
	}{}
	s := d[:]
	d[1].x = 10
	s[2].x = 20
	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
}

func Test_reSlice(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:6] // [2 3 4 5]
	s1[2] = 104

	s2 := s1[2:7] // [104 5 6 7 8]
	s2[1] = 205
	s2[3] = 207

	//s2影響到s1跟s
	fmt.Println(s1) // [2 3 104 205]
	fmt.Println(s)  // [0 1 2 3 104 205 6 207 8 9]
}

func Test_sliceCopy(t *testing.T) {
	s0 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s0[8:] // [8 9] index8以後
	s2 := s0[:5] // [0 1 2 3 4] index不超過5

	copy(s2, s1) // dst:s2, src:s   from s to s2

	fmt.Println(s0)
	fmt.Println(s2)

	s1[1] = -9
	fmt.Println(s0)
	fmt.Println(s2)

	for _, i := range []int{8, 9} {
		fmt.Printf("&s0-%v %p \n", i, &s0[i])
	}
	for index := 0; index < len(s1); index++ {
		fmt.Printf("&s1-%v %p \n", index, &s1[index])
	}
	for _, i := range []int{0, 1} {
		fmt.Printf("&s0-%v %p \n", i, &s0[i])
	}
	for _, i := range []int{0, 1} {
		fmt.Printf("&s2-%v %p \n", i, &s2[i])
	}
}

func Test_memoryArraySlice(t *testing.T) {

	array := [4]int{1, 1, 1, 1}
	slice := array[:2]

	sliceVal(slice, 1, 20)  //clone的slice會影響
	arrayVal(array, 1, -20) //clone的array沒影響

	array[0] = 10 //直接影響slice

	fmt.Print("slice:")
	fmt.Println(slice)
	fmt.Print("array:")
	fmt.Println(array)

	printArrayItemAddress(&array)
	printSliceItemAddress(&slice)

	slice = append(slice, 30)
	slice = append(slice, 40)

	//同array items's address
	printSliceItemAddress(&slice)

	//items's address一樣, 會互相影響
	slice[2] = 31
	array[0] = -10

	//slice apped超過array後,不同於array items's address
	slice = append(slice, 50)
	printSliceItemAddress(&slice)

	//items's address不一樣, 不會互相影響
	slice[3] = 41
	array[1] = -20

	fmt.Print("slice:")
	fmt.Println(slice)
	fmt.Print("array:")
	fmt.Println(array)
}

func Test_memory2Slice(t *testing.T) {

	s1 := []int{1, 1, 1, 1}
	s2 := s1[:2]

	//clone的slice會影響
	sliceVal(s2, 1, 20)

	s1[0] = 10 //直接影響s2

	fmt.Print("s2:")
	fmt.Println(s2)
	fmt.Print("s1:")
	fmt.Println(s1)

	printSliceItemAddress(&s1)
	printSliceItemAddress(&s2)

	s2 = append(s2, 30)
	s2 = append(s2, 40)

	//同s1 items's address
	printSliceItemAddress(&s2)

	//items's address一樣, 會互相影響
	s2[2] = 31
	s1[0] = -10

	//s2 apped超過s1後,不同於s1 items's address
	s2 = append(s2, 50)
	printSliceItemAddress(&s2)

	//items's address不一樣, 不會互相影響
	s2[3] = 41
	s1[1] = -20

	fmt.Print("s2:")
	fmt.Println(s2)
	fmt.Print("s1:")
	fmt.Println(s1)
}
