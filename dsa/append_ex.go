package dsa

import "fmt"

//append
func append_ex1() {
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)
	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s, s2)
}

func append_ex2() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := data[:3]
	s2 := append(s, 100, 200) //
	fmt.Println(data)
	fmt.Println(s)
	fmt.Println(s2)

	fmt.Println(&s[0], &s2[0], &data[0]) //一樣
}

func append_ex3() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	s = append(s, 100, 200) //超過s.cap 重新分配
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0]) //不一樣
}

func append_ex4() { //通常以兩倍容量重新分配底層數組
	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}
