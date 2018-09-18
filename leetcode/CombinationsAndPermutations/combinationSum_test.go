package leetcode

import (
	"fmt"
	"testing"
	"sort"
)

func combinationSum(nums []int, target int) [][]int{
    list := &[][]int{}
	sort.Sort(sort.IntSlice(nums))

    backtrack(list, []int{}, nums, target, 0)
    return *list
}

func backtrack(list *[][]int,  tempList []int, nums []int, remain int, start int){
    if remain < 0 {
		return
	} else if remain == 0 {
		//*list = append(*list, tempList)//直接這樣會影響到tempList
		_tempList := make([]int, len(tempList))
		copy(_tempList, tempList)
		*list = append(*list, _tempList)
	}  else{ 
        for i := start; i < len(nums); i++ {
			tempList = append(tempList, nums[i])
            backtrack(list, tempList, nums, remain - nums[i], i) // not i + 1 because we can reuse same elements
			tempList = tempList[:len(tempList)-1] //移除最後一個			
        }
    }
}

func Test_combinationSum(t *testing.T) {

	// fmt.Println(combinationSum([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum([]int{2, 50}, 200))
	//fmt.Println(combinationSum([]int{1, 2, 3}, 32))
}

func combinationSumMy(candidates []int, target int) [][]int {

	dp := make([][][]int, target+1)

	dp[0] = [][]int{{}}

	length := len(candidates)
	for i := 0; i < length; i++ {
		candidate := candidates[i]
		for j := candidate; j <= target; j++ {
			if len(dp[j-candidate]) > 0 {

				for _, nums := range dp[j-candidate] {
					_length := len(nums)

					if _length == 0 {
						dp[j] = append(dp[j], []int{candidate})
						break
					}

					_nums := make([]int, _length)
					copy(_nums, nums)

					_nums = append(_nums, candidate)

					dp[j] = append(dp[j], _nums)
				}

			}
		}
	}

	return dp[target]
}

func Test_combinationSumMy(t *testing.T) {

	//fmt.Println(combinationSum([]int{1, 2, 3}, 4))
	fmt.Println(combinationSumMy([]int{2, 50}, 200))
	//fmt.Println(combinationSum([]int{1, 2, 3}, 32))
}
