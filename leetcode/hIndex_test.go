package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func hIndex(citations []int) int {

	sort.Ints(citations)

	count := len(citations)

	for i := 0; i < count; i++ {
		if citations[i] >= count-i {
			return count - i
		}
	}

	return 0
}

func Test_hIndex(t *testing.T) {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
