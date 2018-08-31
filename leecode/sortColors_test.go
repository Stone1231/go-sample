package leecode

import (
	"fmt"
	"testing"
)

func sortColors(nums *[]int) {
	s1, s2, s3 := []int{},[]int{},[]int{}
	length := len(*nums)
	for i := 0; i < length; i++ {
		switch (*nums)[i] {
		case 0:
			s1 = append(s1,0)
		case 1:
			s2 = append(s2,1)
		case 2:
			s3 = append(s3,2)
		}
	}

	*nums = append(append(s1,s2...),s3...)
}

func Test_sortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(&nums)
	fmt.Println(nums)
}
