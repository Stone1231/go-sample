package examex

import (
	"fmt"
	"testing"
)

//Stacking defers
//https://tour.golang.org/flowcontrol/13
func Test_deferMuilt(t *testing.T) {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
