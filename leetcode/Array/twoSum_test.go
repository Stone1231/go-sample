package leetcode

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {

	length := len(nums)

	m := map[int]int{}

	for i := 0; i < length; i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}

	return make([]int, 2)
}

func brute_force(nums []int, target int) []int {

	length := len(nums)

	res := make([]int, 2)

	end := false

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				end = true
				break
			}
		}
		if end {
			break
		}
	}

	return res
}

func Test_twoSum(t *testing.T) {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 5))
}
