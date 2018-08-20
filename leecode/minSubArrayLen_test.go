package leecode

import (
	"fmt"
	"math"
	"testing"
)

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	min_len := math.MaxUint32
	left := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= s {

			if min_len > i+1-left {
				min_len = i + 1 - left
			}

			sum -= nums[left]
			left++
		}
	}

	if min_len != math.MaxUint32 {
		return min_len
	}
	return 0
}

func Test_minSubArrayLen(t *testing.T) {
	fmt.Println(minSubArrayLen(213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}))
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(100, []int{}))
}
