package leetcode

import (
	"fmt"
	"testing"
)


func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i+num <= target {
				//記錄加總到i+num這個數字的總次數
				//目前數字i+num
				//上個數字為i
				dp[i+num] += dp[i]
				fmt.Printf("dp[%v+%v] += dp[%v] = %v \n", i, num, i, dp[i+num])
			}
		}
	}

	return dp[target]
}

func Test_combinationSum4(t *testing.T) {

	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	//fmt.Println(combinationSum4([]int{1, 50}, 200))
	//fmt.Println(combinationSum4([]int{1, 2, 3}, 32))
}

//只能取得Group組合數
func combinationSum4Group(nums []int, target int) int {
	dp := make([]int, target+1)
	length := len(nums)
	dp[0] = 1

	for i := 0; i < length; i++ {
		num := nums[i]

		for j := num; j <= target; j++ {
			dp[j] += dp[j-num]
			fmt.Printf("dp[%v] += dp[%v-%v] = %v \n", j, j, num, dp[j])
		}

		//同上
		// for j := 0; j <= target; j++ {
		// 	if j+num <= target {
		// 		dp[j+num] += dp[j]
		// 		fmt.Printf("dp[%v+%v] += dp[%v] = %v \n", j, num, j, dp[j+num])
		// 	}
		// }
	}

	return dp[target]
}

func Test_combinationSum4Group(t *testing.T) {
	fmt.Println(combinationSum4Group([]int{1, 2, 3}, 4))
	// fmt.Println(combinationSum4Group([]int{1, 50}, 200))
	// fmt.Println(combinationSum4Group([]int{1, 2, 3}, 32))	
}

func combinationSum4Job(dp *[]int, current int, nums *[]int, target int) {
	(*dp)[current]++
	for _, _num := range *nums {
		next := current + _num
		if next <= target{
			combinationSum4Job(dp,next, nums,target)
		}	
	}
}

func combinationSum4MemoryError3(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	combinationSum4Job(&dp,0,&nums,target)

	return dp[target]
}

func Test_combinationSum4MemoryError3(t *testing.T) {
	fmt.Println(combinationSum4MemoryError3([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum4MemoryError3([]int{1, 50}, 200))
	fmt.Println(combinationSum4MemoryError3([]int{1, 2, 3}, 32))	
}

func combinationSum4MemoryError2(nums []int, target int) int {

	dp := make([][][]int, target+1)
	length := len(nums)
	dp[0] = [][]int{}
	//dp[0][0] = make([]int, length)
	dp[0] = append(dp[0], make([]int, length))

	for i := 0; i < length; i++ {
		num := nums[i]
		for j := num; j <= target; j++ {
			prev := dp[j-num]
			if len(prev) > 0 {

				_index := len(dp[j]) - 1
				_ss := make([][]int, len(prev))
				dp[j] = append(dp[j], _ss...)

				for _, value := range prev {
					_index++
					dp[j][_index] = make([]int, len(value))
					copy(dp[j][_index], value)
					dp[j][_index][i]++
				}

			}
		}
	}
	if len(dp[target]) == 0 {
		return 0
	}

	ss := dp[target]
	sum := 0
	for i := 0; i < len(ss); i++ {
		s := ss[i]
		v1 := 0
		v2 := 1
		for j := 0; j < len(s); j++ {
			v1 += s[j]
			_s := s[j]
			for _s > 0 {
				v2 *= _s
				_s--
			}
		}
		_sum := 1
		for v1 > 1 {
			_sum *= v1
			v1--
		}
		_sum /= v2
		sum += _sum
	}

	return sum
}

func Test_combinationSum4MemoryError2(t *testing.T) {

	//fmt.Println(combinationSum4MemoryError2([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum4MemoryError2([]int{1, 50}, 200))
	//fmt.Println(combinationSum4MemoryError2([]int{1, 2, 3}, 32))
}

func combinationSum4MemoryError(nums []int, target int) int {

	length := len(nums)
	res := 0
	dp0 := &[]int{0}

	for len(*dp0) > 0 {
		dpn := []int{}
		for i := 0; i < length; i++ {
			count := len(*dp0)
			for j := 0; j < count; j++ {
				sum := (*dp0)[j] + nums[i]
				if sum < target {
					dpn = append(dpn, sum)
				}

				if sum == target {
					res++
				}
			}
		}
		dp0 = &dpn
	}

	return res
}

func Test_combinationSum4MemoryError(t *testing.T) {

	fmt.Println(combinationSum4MemoryError([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum4MemoryError([]int{1, 50}, 200))
	fmt.Println(combinationSum4MemoryError([]int{1, 2, 3}, 32))
}
