//https://leetcode.com/problems/house-robber-ii/description/
package leetcode

import (
	"fmt"
	"testing"
)

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	res1 := robMax(nums, 0, length-2)
	res2 := robMax(nums, 1, length-1)

	if res1 > res2 {
		return res1
	}
	return res2
}

//refrence from https://leetcode.com/problems/house-robber-ii/discuss/59934/Simple-AC-solution-in-Java-in-O(n)-with-explanation
func robMax(num []int, start, end int) int {
	include, exclude := 0, 0 //每個點都有include,exclude兩種max總和
	is, es := []int{}, []int{}
	for j := start; j <= end; j++ {

		current := exclude + num[j] //到上個exclude的總和再加上目前

		if exclude < include { //上個include跟上個exclude 看那個大
			exclude = include
		}

		include = current //目前include

		is = append(is, include)
		es = append(es, exclude)
	}

	fmt.Printf("  num[]: %v \n", num)
	fmt.Printf("include: %v \n", is)
	fmt.Printf("exclude: %v \n", es)

	if include > exclude {
		return include
	}
	return exclude
}

func Test_rob(t *testing.T) {
	//fmt.Println(rob([]int{1, 20, 300, 4000, 50000, 600000, 7000000}))
	fmt.Println(rob([]int{1, 20, 300, 400000, 50000, 6000, 7000000, 80000000, 900000000}))
}

var max int

func robMy(nums []int) int {
	count := len(nums)

	total := count
	if count > 3 {
		total = 3
	}

	max = 0

	for i := 0; i < total; i++ {
		m := map[int]int{}
		m[i] = nums[i]

		next(&nums, i, i, count, m)
	}

	return max
}

func next(nums *[]int, start, prev, count int, m map[int]int) {

	sum := m[prev]

	if prev > count-4 {
		if sum > max {
			max = sum
		}
		// return
	}

	prev += 2 //空一格後的下一個
	if start == 0 && prev == count-1 {
		return
	}
	if prev < count {
		_sum := sum + (*nums)[prev]

		if _sum > m[prev] {
			m[prev] = _sum
			next(nums, start, prev, count, m)
		}
	}

	prev++ //空兩格後的下一個
	if start == 0 && prev == count-1 {
		return
	}
	if prev < count {
		_sum := sum + (*nums)[prev]

		if _sum > m[prev] {
			m[prev] = _sum
			next(nums, start, prev, count, m)
		}
	}
}

func Test_robMy(t *testing.T) {
	// fmt.Println(robMy([]int{2, 3, 2}))
	// fmt.Println(robMy([]int{1, 2, 3, 1}))
	// fmt.Println(robMy([]int{}))
	// fmt.Println(robMy([]int{2, 7, 9, 3, 1}))
	// fmt.Println(robMy([]int{1, 1, 3, 6, 7, 10, 7, 1, 8, 5, 9, 1, 4, 4, 3}))

	fmt.Println(robMy([]int{1, 20, 300, 4000, 50000, 600000, 7000000}))
	fmt.Println(robMy([]int{1, 20, 300, 400000, 50000, 6000, 7000000, 80000000, 900000000}))
}
