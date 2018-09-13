package leetcode

import (
	"fmt"
	"testing"
)

func permute(nums []int) [][]int {

	count := len(nums)
	aa := [][]int{}
	aa = append(aa, []int{})

	//check contains
	cc := [][]bool{}
	cc = append(cc, make([]bool, count))

	_count := count

	for _count > 0 {
		_aa := [][]int{}
		_cc := [][]bool{}
		for i := 0; i < count; i++ {
			for j, a := range aa {
				if !cc[j][i]{
					_c := make([]bool, len(cc[j]))
					copy(_c, cc[j])
					_c[i] = true
					_cc = append(_cc, _c)

					_a := make([]int, len(a))
					copy(_a, a)
					_a = append(_a, nums[i])
					_aa = append(_aa, _a)
				}
			}
		}
		aa = _aa
		cc = _cc
		_count--
	}

	return aa
}

func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}
