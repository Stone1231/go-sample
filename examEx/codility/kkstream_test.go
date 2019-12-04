package examex

import (
	"fmt"
	"testing"
	"math"
)

func Solution1(predicted []float64, observed []float64) float64 {
	count := len(predicted)
	total := float64(0)
	for i := 0; i < count; i++ {
		total += math.Pow(predicted[i] - observed[i], 2)		
	}
    return math.Sqrt(total/float64(count))
}

func Test_exam1(t *testing.T) {
	fmt.Println(Solution1([]float64{4, 25,  0.75, 11}, []float64{3, 21, -1.25, 13}))
}

type Tree struct {
	X       int
	L *Tree
	R *Tree
}

var sum = 0

func sumNode(T *Tree, V int) {
	if (T == nil){
		return
	}

	if (T.X >= V){
		sum++
		sumNode(T.L, T.X)
		sumNode(T.R, T.X)
	} else{
		sumNode(T.L, V)
		sumNode(T.R, V)
	}
}

func Solution2(T *Tree) int {

	if (T != nil){
		sum = 1
	}

	if (T.L != nil){
		sumNode(T.L, T.X)
	}

	if (T.R != nil){
		sumNode(T.R, T.X)
	}	

	return sum
}

func Test_exam2(t *testing.T) {
	// fmt.Println(Solution(Tree(8, (2, (8, nil, nil), (7, nil, nil)), (6, nil, nil)))))
	newRoot := &Tree{
		X: 1,
		L: nil,
		R: nil,
	}
	fmt.Println(Solution2(newRoot))
}

func Solution(T []int) int {
	length := len(T)
	l := 0
	r := length -1
	max_l := T[l]
	min_r := T[r]

	for {
		if (l+1 < length -1 && T[l+1] < min_r){
			if (T[l+1] > max_l){
				max_l = T[l+1]
			}
			l++			
		}
		
		if (r-1 >0 && T[r-1] > max_l){
			if (T[r-1] < min_r){
				min_r = T[r-1]
			}
			r--			
		}

		if (l == r-1 || l ==r){
			break;
		}
	}


    return l + 1
}

func Test_exam(t *testing.T) {
	fmt.Println(Solution([]int{-5, -5, -5, -42, 6, 12}))
}