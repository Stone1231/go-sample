package leetcode

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)//value:記錄加總到index這個數字的最小次數
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	//dp[0] = 0 def value, 略

	length := len(coins)
	for i := 0; i < length; i++ {
		coin := coins[i]
		for j := coin; j <= amount; j++ {
			//記錄加總到j這個數字的次數
			//目前數字j
			//上個數字為j-coin
			count := dp[j-coin] + 1
			if count < dp[j] { //最少次數
				dp[j] = count

				// if j%coin == 0 {
				// 	fmt.Printf("dp[%v] = %v \n", j, count)
				// } else {
				// 	fmt.Printf("dp[%v]*= %v ", j, count)
				// 	fmt.Printf("dp[%v - %v] + 1 = ", j, coin)
				// 	fmt.Printf("dp[%v] + 1 = v = %v \n", j-coin, count)
				// }
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func Test_coinChange(t *testing.T) {

	//fmt.Println(coinChange([]int{1, 2, 5}, 11))
	//fmt.Println(coinChange([]int{2}, 3))
	//fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	//fmt.Println(coinChange([]int{2, 4, 5}, 39))
	fmt.Println(coinChange([]int{288,160,10,249,40,77,314,429}, 6249))

}

func coinChangeMy2(coins []int, amount int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	count := len(coins)
	m := make([]int, count)
	start := 0
	last := 0

	r := -1
	t := 0

	for {
		for i := start; i < count; i++ {
			coin := coins[i]
			m[i] = 0
			if amount >= coin {
				m[i] = amount / coin
				amount = amount % coin
				last = i
			}
		}

		_r := 0
		for _, v := range m {
			_r += v
		}
		t++
		fmt.Printf("%v %v %v \n", t, _r, m)

		if amount == 0 {
			// _r := 0
			// for _, v := range m {
			// 	_r += v
			// }

			if r == -1 || r > _r {
				r = _r
			}

			// fmt.Print(_r)
			// fmt.Println(m)
		}

		j := last

		if j > 0 && j == count-1 {
			amount += coins[j] * m[j]
			j--
		}

		for j > 0 && m[j] == 0 {
			j--
		}

		if j == 0 && m[j] == 0 {
			break
		}

		amount += coins[j]
		m[j]--

		start = j + 1
	}

	return r
}

func coinChangeMy(coins []int, amount int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	r := 0

	for _, coin := range coins {
		if amount >= coin {
			r += amount / coin
			amount = amount % coin
		}
	}

	if amount > 0 {
		r = -1
	}

	return r
}


