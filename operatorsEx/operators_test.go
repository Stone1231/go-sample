package operatorsEx

import (
	"fmt"
	"testing"
	"math"
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

	//Get Last Set Bit	
	//1  =                                    1
	//-1 =     11111111111111111111111111111111
	//1 & -1 =                                1
	fmt.Println(1 & -1)
	//2  =                                   10
	//-2 =     11111111111111111111111111111110
	//2 & -2 =                               10
	fmt.Println(2 & -2)

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

//二補數（2's complement）
func Test_2complement(t *testing.T) {

	num := 2

	//-num
	negNum := math.MaxUint32 + num + 1
	fmt.Printf(" num = %b \n", num)
	fmt.Printf("-num = %b \n", negNum)
}

func Test_operatorsRem(t *testing.T) {
	//75除32的餘數
	//指數運算 只能用2的次方
	//ex 除以2的5次方32
	fmt.Println(75 % 32)
	fmt.Println(75 & 31)       //最大餘數31 &
	fmt.Println(75 - 75>>5<<5) //位元的左右移,2的5次方
}

func Test_operatorsQuotient(t *testing.T) {
	//75除32的商
	//指數運算 只能用2的次方
	//ex 除以2的5次方32
	fmt.Println(75 >> 5) //位元的左右移,2的5次方
	fmt.Println(75 / 32) //同上	
}

func Test_operatorsXOR(t *testing.T) {

	fmt.Println(^5)    //-6
	fmt.Println(5 ^ 0) // 5

	ans := 0
	//                             1, 2, 3, 1, 3, 2互相抵消
	for _, value := range []int{7, 1, 2, 3, 1, 3, 2} {
		ans ^= value
		fmt.Printf("%v->", ans)
	}

}

//取右邊第一個1
func Test_getLastSetBit(t *testing.T) {
	for i := 0; i < 100; i++ {
		
		fmt.Printf("%v:%v", i, i & -i)
		fmt.Printf(" (%b & %b) \n", i, math.MaxUint32-i + 1)//才能顯示實際的二進位數
	}
}

func Test_bitNegative(t *testing.T) {
	fmt.Printf("%b\n", -7) //-111
	fmt.Printf("\n%b\n", 3)	//                                               11
	fmt.Printf("%b\n", math.MaxUint32-7 + 1) //11111111111111111111111111111001
	fmt.Printf("%b\n", 3 & -7) //                                             1
}
