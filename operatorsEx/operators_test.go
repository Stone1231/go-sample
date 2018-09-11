package operatorsEx

import (
	"fmt"
	"testing"
)

func Test_operators(t *testing.T) {
	// Use bitwise OR | to get the bits that are in 1 OR 2
	// 1     = 00000001
	// 2     = 00000010
	// 1 | 2 = 00000011 = 3
	fmt.Println(1 | 2)

	// Use bitwise OR | to get the bits that are in 1 OR 5
	// 1     = 00000001
	// 5     = 00000101
	// 1 | 5 = 00000101 = 5
	fmt.Println(1 | 5)

	// Use bitwise XOR ^ to get the bits that are in 3 OR 6 BUT NOT BOTH
	// 3     = 00000011
	// 6     = 00000110
	// 3 ^ 6 = 00000101 = 5
	fmt.Println(3 ^ 6)

	// Use bitwise AND & to get the bits that are in 3 AND 6
	// 3     = 00000011
	// 6     = 00000110
	// 3 & 6 = 00000010 = 2
	fmt.Println(3 & 6)

	// Use bit clear AND NOT &^ to get the bits that are in 3 AND NOT 6 (order matters)
	// 3      = 00000011
	// 6      = 00000110
	// 3 &^ 6 = 00000001 = 1
	fmt.Println(3 &^ 6)

	var A = 60          //0011 1100
	fmt.Println(A << 2) //A << 2 will give 240 which is 1111 0000
	fmt.Println(A >> 2) //A >> 2 will give  15 which is 0000 1111
	fmt.Println(A >> 1)

	fmt.Println(^5)
	// 5  00000101
	//-6  11111010 (one's complement of 5, the value was -6)	
}

func Test_operatorsRem(t *testing.T) {
	//75除32的餘數
	//指數運算 只能用2的次方
	//ex 除以2的5次方32
	fmt.Println(75 % 32)
	fmt.Println(75 & 31)//最大餘數31 &
	fmt.Println(75 - 75 >> 5 << 5) //位元的左右移,2的5次方
}

func Test_operatorsQuotient(t *testing.T) {
	//75除32的商
	//指數運算 只能用2的次方
	//ex 除以2的5次方32
	fmt.Println(75 / 32)
	fmt.Println(75 >> 5) //位元的左右移,2的5次方

	fmt.Println(3 ^ 0)
}
