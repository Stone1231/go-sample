package leetcode

import (
	"fmt"
	"testing"
)

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	count := len(nums)
	
	//Use bitwise XOR ^ to get the bits that are in 
	//a ^ b
	//a OR b BUT NOT BOTH

	//Use bit clear AND NOT &^ to get the bits that are in 
	//a &^ b
	//a AND NOT b (order matters)

	for i := 0; i < count; i++ {
		// ones = (ones ^ nums[i]) &^ twos
		// twos = (twos ^ nums[i]) &^ ones

		//同上, 上面是縮寫
		ones ^= nums[i] //a OR b BUT NOT BOTH
		ones &^= twos //a AND NOT b

		twos ^= nums[i]
		twos &^= ones
	}

    return ones
}

func singleNumberMy(nums []int) int {
	m := map[int]int{}
	ans := -1

	for _, num := range nums {
		m[num]++
		if m[num] == 3 {
			delete(m, num)
		}
	}

	for k, v := range m {
		if v == 1 {
			ans = k
			break
		}
	}

	return ans
}

func Test_singleNumber(t *testing.T) {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber([]int{43, 16, 45, 89, 45, -2147483648, 45, 2147483646, -2147483647, -2147483648, 43, 2147483647, -2147483646, -2147483648, 89, -2147483646, 89, -2147483646, -2147483647, 2147483646, -2147483647, 16, 16, 2147483646, 43}))
}
