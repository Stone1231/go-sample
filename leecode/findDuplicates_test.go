package leecode

import (
	"fmt"
	"math"
	"testing"
)

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

//Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array),
//some elements appear twice and others appear once.
func findDuplicates(nums []int) []int {

	_nums := &[]int{}

	for _, num := range nums {

		if nums[abs(num)-1] < 0 {
			*_nums = append(*_nums, abs(num))
		} else {
			nums[abs(num)-1] *= -1
		}

	}

	fmt.Println(nums)

	return *_nums
}

func findDuplicatesMy(nums []int) []int {

	m := map[int]int{}

	_nums := []int{}

	for _, num := range nums {
		if m[num] > 0 {
			_nums = append(_nums, num)
		}
		m[num]++
	}

	return _nums
}

func Test_findDuplicates(t *testing.T) {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func Test_findDuplicatesMy(t *testing.T) {
	fmt.Println(findDuplicatesMy([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
