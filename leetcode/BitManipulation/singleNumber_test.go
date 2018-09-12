package leetcode

import (
	"fmt"
	"testing"
)

//裡面的元素每個都出現了兩次，除了一個特殊的
func singleNumberBase(nums []int) int {
	ans := 0
	count := len(nums)
	for i := 0; i < count; i++ {
		ans ^= nums[i]		
	}
	return ans 
}

func Test_singleNumberBase(t *testing.T) {
	fmt.Println(singleNumberBase([]int{1, 2, 3, 1, 2}))
}