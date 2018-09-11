package leetcode

import (
	"fmt"
	"testing"
	"math"
)

func toHex(num int) string{

	if num == 0 {
		return "0"
	}
	
	arr := "0123456789abcdef"
	ans := ""

	//用二補數表示負數
	if num < 0 {
		num = math.MaxUint32 + num + 1
	}
	
	for num > 0 {
	  ans = string(arr[num & 15]) + ans
	  num = num >> 4
	}
	
	return ans
}

func toHexMy(num int) string {

	if num == 0 {
		return "0"
	}

	arr := "0123456789abcdef"

	hex := ""

	if num < 0 {
		num = math.MaxUint32 + num + 1
	}	

	for num != 0 {
		hex = string(arr[num % 16]) + hex
		num = num / 16
	}

	return hex
}

func Test_toHex(t *testing.T) {

	fmt.Println(toHex(26))
	fmt.Println(toHex(-1))
	fmt.Println(toHex(-2))

	fmt.Println(toHexMy(26))
	fmt.Println(toHexMy(-1))
	fmt.Println(toHexMy(-2))

}
