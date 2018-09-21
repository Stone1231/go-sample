package leetcode

import (
	"fmt"
	. "sample/converex"
	"sort"
	"strconv"
	"testing"
)

func largestNumber(nums []int) string {

	sort.SliceStable(
		nums,
		func(i, j int) bool {
			ijStr := fmt.Sprintf("%v%v", nums[i], nums[j])
			jiStr := fmt.Sprintf("%v%v", nums[j], nums[i])
			return ijStr > jiStr
		})

	if len(nums) > 0 && nums[0] == 0 {
		return "0"
	}

	return ArrayToString(nums, "")
}

func Test_largestNumber(t *testing.T) {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{0, 0}))
}

func largestNumberMy(nums []int) string {
	return ArrayToString(customMergeSort(nums), "")
}

func customMergeSort(unsorted []int) []int {
	if len(unsorted) <= 1 {
		return unsorted
	}

	middle := len(unsorted) / 2
	left := unsorted[:middle]
	right := unsorted[middle:]

	left = customMergeSort(left)
	right = customMergeSort(right)

	return customMerge(left, right)
}

func customMerge(left []int, right []int) []int {
	res := []int{}

	l := 0
	r := 0
	lLast := len(left) - 1
	rLast := len(right) - 1

	for l <= lLast && r <= rLast {
		if strconv.Itoa(left[l]) > strconv.Itoa(right[r]) {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}
	if l > lLast {
		res = append(res, right[r:]...)
	} else {
		res = append(res, left[l:]...)
	}
	return res
}

func Test_largestNumberMy(t *testing.T) {
	fmt.Println(largestNumberMy([]int{3, 30, 34, 5, 9}))
}

func largestNumberMyError(nums []int) string {

	ints := []int{}

	for _, num := range nums {
		for num > 0 {
			v := num % 10
			ints = append(ints, v)
			num = num / 10
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

	return ArrayToString(ints, "")
}

func Test_largestNumberMyError(t *testing.T) {
	fmt.Println(largestNumberMyError([]int{3, 30, 34, 5, 9}))
}
