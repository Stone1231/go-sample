package dsa

import "fmt"

func slice_ex1() {

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(data[1:4:5]) // [low : high : max]
	s := data[1:4:5]
	fmt.Println(s, len(s), cap(s))

	for index := 0; index < len(s); index++ {
		fmt.Println(index, s[index])
	}
	// for var v:= range s {
	// 	fmt.Println(v)
	// }

	fmt.Println(data[:6:8])
	fmt.Println(data[5:])
	fmt.Println(data[:3])
	fmt.Println(data[:])
}

func slice_ex2() {
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	fmt.Println(data)
}

func slice_ex3() {
	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]int, 6, 8)
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 6)
	fmt.Println(s3, len(s3), cap(s3))
}

func slice_ex4() {
	s := []int{0, 1, 2, 3}

	p := &s[2] // *int
	*p += 100
	fmt.Println(s)
}

func slice_ex5() {
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

///reslice
func reslice_ex() {
	// s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s1 := s[2:5]    //[2 3 4]
	// s2 := s1[2:6:7] //[4 5 6 7]
	// s3 := s2[3:6] //error

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5] // [2 3 4]
	s1[2] = 100
	s2 := s1[2:6] // [100 5 6 7]
	s2[3] = 200
	fmt.Println(s)
}

func slice_copy() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:] //index8個以後
	fmt.Println(s)
	s2 := data[:5] //index不超過5
	fmt.Println(s2)
	copy(s2, s) // dst:s2, src:s   from s to s2
	fmt.Println(s2)
	fmt.Println(data)
}
