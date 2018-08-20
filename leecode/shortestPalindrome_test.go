package leecode

import (
	"bytes"
	"fmt"
	"testing"
)

func shortestPalindrome3(s string) string {
	length := len(s)

	if length == 0 {
		return ""
	}

	desc_i := length - 1
	i := 0
	right := 0
	right_jump := 0

	for i < desc_i {
		right += right_jump
		desc_i = length - 1 - right
		i = 0

		right_jump = 1
		jump := 0

		for i < desc_i && s[i] == s[desc_i] {

			if s[0] == s[i] && i > 0 {
				if jump == 0 {
					jump = i
					right_jump = i
				}
				if i-right_jump != jump {
					right_jump = i
				}
			}

			desc_i--
			i++
		}
	}

	var buffer bytes.Buffer
	for index := 0; index < right; index++ {
		buffer.WriteByte(s[length-1-index])
	}

	buffer.WriteString(s)
	return buffer.String()
}

func shortestPalindrome2(s string) string {
	n := len(s)
	right := 0
	i := 0
	var start, end int

	for i < (n+1)/2 {
		start = i
		for i < n && s[i] == s[start] {
			i++
		}
		end = i - 1
		for start-1 >= 0 &&
			end+1 < n &&
			s[start-1] == s[end+1] {
			start--
			end++
		}
		if start == 0 && end > right {
			right = end
		}
	}

	var buffer bytes.Buffer
	if right < n-1 {
		for index := n - 1; index > right; index-- {
			buffer.WriteByte(s[index])
		}
	}

	return buffer.String() + s
}

func shortestPalindrome(s string) string {

	s_length := len(s)

	rev_s := reverseStr(s, s_length)

	l := s + "#" + rev_s
	l_length := len(l)

	p := failure(l)

	return rev_s[:s_length-p[l_length-1]] + s
}

func failure(s string) []int {
	length := len(s)

	p := make([]int, length)

	for i := 1; i < length; i++ {
		p_len := p[i-1] //i-1 的前綴長度, 也是前綴下個index剛好比對s[i]
		for p_len > 0 && s[i] != s[p_len] {
			p_len = p[p_len-1] //不相同就尋找上個前綴
		}

		if s[i] == s[p_len] {
			p_len++
		}
		p[i] = p_len
	}

	return p
}

func reverseStr(s string, l int) string {
	var buffer bytes.Buffer
	for index := l - 1; index >= 0; index-- {
		buffer.WriteByte(s[index])
	}
	return buffer.String()
}

func Test_shortestPalindrome(t *testing.T) {
	fmt.Println(shortestPalindrome("aaabbbccc"))
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
	fmt.Println(shortestPalindrome("abcbae"))
	fmt.Println(shortestPalindrome("a"))
	fmt.Println(shortestPalindrome("abbacd"))

	fmt.Println(failure("113113211131131"))
	fmt.Println(failure("113113211131132"))
}
