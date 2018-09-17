package leetcode

import (
	"fmt"
	"testing"
)

func dfs(start, n, k int, curArr []int, result *[][]int) {

	if len(curArr) == k {
		_curArr := make([]int, k)
		copy(_curArr, curArr)
		*result = append(*result, _curArr)
		return
	}

	for i := start; i <= n; i++ {
		curArr = append(curArr, i)
		dfs(i+1, n, k, curArr, result)
		curArr = curArr[:len(curArr)-1] //移除最後一個
	}
}

func combine(n int, k int) [][]int {
	result := &[][]int{}

	dfs(1, n, k, []int{}, result)

	return *result
}

func Test_combine(t *testing.T) {

	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
	fmt.Println(combine(2, 1))
	fmt.Println(combine(3, 3))
}

func combinePoor(n int, k int) [][]int {
	//aa := [][]int{}
	//aa = append(aa, []int{1})
	aa := make([][]int, n)
	//a := make([]int,k)

	for i := 0; i < n; i++ {
		//aa = append(aa, []int{i})
		aa[i] = []int{i + 1}
	}

	i := 0
	for k-1 > i {

		_aa := [][]int{}

		for _, a := range aa {
			for j := a[i]; j < n; j++ {
				// _a := make([]int, len(a))
				// copy(_a, a)
				// _a = append(_a, j+1)
				_a := make([]int, k)
				copy(_a, a)
				_a[i+1] = j + 1

				_aa = append(_aa, _a)
			}
		}

		i++

		aa = _aa
	}

	return aa
}

func Test_combinePoor(t *testing.T) {

	fmt.Println(combinePoor(4, 2))
	// fmt.Println(combinePoor(1, 1))
	// fmt.Println(combinePoor(2, 1))
	// fmt.Println(combinePoor(3, 3))
}

