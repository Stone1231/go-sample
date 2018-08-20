package leecode

import (
	"fmt"
	"strings"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("c"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("123a567a1238"))
}

//Assuming ASCII 128
func lengthOfLongestSubstring(s string) int {

	l := len(s)
	//m := map[byte]int{}
	m := [128]int{} //效能較map更好!!
	count := 0
	start := 0

	for index := 0; index < l; index++ {
		c := s[index]
		if m[c] > start {
			_count := index - start

			if _count > count {
				count = _count
			}
			start = m[c]
		}
		m[c] = index + 1
	}

	_count := l - start
	if _count > count {
		count = _count
	}

	return count
}

func lengthOfLongestSubstring_bad(s string) int {

	m := map[byte]int{}

	ms := ""
	_s := ""

	for index := 0; index < len(s); index++ {
		c := s[index]
		if m[c] > 0 {

			if len(_s) > len(ms) {
				ms = _s
			}

			strs := strings.Split(_s, string(c))
			_s = strs[1]

			m = map[byte]int{}
			for j := 0; j < len(_s); j++ {
				m[_s[j]] = 1
			}
		}
		_s += string(c)
		m[c]++
	}

	if len(_s) > len(ms) {
		ms = _s
	}

	return len(ms)
}
