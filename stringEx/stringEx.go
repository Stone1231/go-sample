package main

import (
	"fmt"
)

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
}

func test1() {
	s := "abc"
	fmt.Println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
}

func test2() {
	s := `a
b\r\n\x00
c`
	fmt.Println(s)
}

func test3() {
	s := "Hello, " +
		"World!"
	fmt.Println(s)
}

func test4() {
	s := "Hello, World!"
	s1 := s[:5]
	s2 := s[7:]
	s3 := s[1:5]

	fmt.Println(s1, s2, s3)
	// Hello
	// World!
	// ello
}

func test5() {
	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '們'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac")
}

func test6() {
	s := "abcd"
	bs := []byte(s)
	bs[1] = 'B'
	println(string(bs))

	u := "電腦"
	us := []rune(u)
	us[1] = '話'
	println(string(us))
}

func test7() {

	s := "abc漢字"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c,", s[i]) // byte
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%c,", r) // rune
	}
	fmt.Println()
}
