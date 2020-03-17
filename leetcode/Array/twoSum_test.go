package leetcode

import (
	"fmt"
	"testing"
)

// 一個陣列和一個數字，數字會是這個陣列裡的某兩個數字的總和，答案就是這兩個數字在陣列的位置

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

func bruteForce(nums []int, target int) []int {

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
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 6))
}
