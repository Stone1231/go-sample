package stringEx

import (
	"fmt"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	s := "abc"
	_s := reverseStr(s)
	if _s == "cba" {
		t.Log("reverseStr PASS")
	} else {
		t.Error("reverseStr FAIL")
	}
}

func Test_strChar(t *testing.T) {
	s := "abc"
	fmt.Println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
}

func Test_textStr(t *testing.T) {
	s := `a
 b\r\n\x00
c`
	fmt.Println(s)
}

func Test_addStr(t *testing.T) {
	s := "Hello, " +
		"World!"
	fmt.Println(s)
}

func Test_subStr(t *testing.T) {
	s := "Hello, World!"
	s1 := s[:5]
	s2 := s[7:]
	s3 := s[1:5]

	fmt.Println(s1, s2, s3)
	// Hello
	// World!
	// ello
}

func Test_rune(t *testing.T) {
	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '們'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac")
}

func Test_byteRune(t *testing.T) {
	s := "abcd"
	bs := []byte(s)
	bs[1] = 'B'
	println(string(bs))

	u := "電腦"
	us := []rune(u)
	us[1] = '話' //超過byte,不能用byte
	println(string(us))
}

func Test_byteRuneStr(t *testing.T) {

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
