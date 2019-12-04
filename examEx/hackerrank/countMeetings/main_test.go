package main

import (
	"fmt"
	"testing"
)

func countMeetings(arrival []int32, departure []int32) int32 {
	// Write your code here
	count := len(arrival)

	var min int32
	var max int32

	for i := 0; i < count; i++ {
		if i == 0 {
			min = arrival[i]
			max = departure[i]
		} else{
			if (arrival[i] < min && departure[i] < min) || (arrival[i] > max && departure[i] > max){
				
			}

		}


	}

	return 0
}

func Test_exam(t *testing.T) {
	fmt.Println(countMeetings([]int32{}, []int32{}))
}
