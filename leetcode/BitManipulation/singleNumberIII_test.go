package leetcode

import (
	"fmt"
	"testing"
)

func singleNumberIII(nums []int) []int {

	// Pass 1 :
	// Get the XOR of the two numbers we need to find
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	// Get its last set bit
	//二進位取右邊最後一個1, 兩數相加一定一個1,一個0	
	diff &= -diff

	// Pass 2 :
	rets := []int{0, 0} // this array stores the two numbers we will return
	for _, num := range nums {
		if (num & diff) == 0 { // the bit is not set
			                   // num沒有1,重覆會抵消,會剩下沒有1的某數
			rets[0] ^= num
		} else { // the bit is set
			     // num有1,重覆會抵消,會剩下有1的某數
			rets[1] ^= num
		}
	}
	return rets
}

func Test_singleNumberIII(t *testing.T) {
	fmt.Println(singleNumberIII([]int{1, 2, 1, 2, 3, 5}))
}
