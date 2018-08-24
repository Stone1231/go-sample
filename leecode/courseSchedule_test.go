package leecode

import (
	"testing"
	"fmt"
)

func checkCycle(
	before int,
	after int,
	m *(map[int]map[int]bool)) bool {

	if ma, ok := (*m)[after]; ok {
		if _, ok := ma[before]; ok {
			return false
		}

		for k := range ma {
			res := checkCycle(before, k, m)
			if !res {
				return false
			}
		}
	}

	return true
}

func canFinishMy(
	numCourses int,
	prerequisites [][]int) bool {

	m := map[int](map[int]bool){}

	courses := map[int]bool{}
	sum := 0

	if len(prerequisites) == 0 {
		return true
	}

	for _, item := range prerequisites {
		after := item[0]
		before := item[1]
		res := checkCycle(after, before, &m)

		if !res {
			return false
		}

		if _, ok := m[after]; !ok {
			m[after] = map[int]bool{}
		}
		m[after][before] = true

		if _, ok := courses[after]; !ok {
			courses[after] = true
			sum++
		}
		if _, ok := courses[before]; !ok {
			courses[before] = true
			sum++
		}

	}

	return numCourses >= sum
}

func Test_canFinishMy(t *testing.T) {
	println(canFinishMy(2, [][]int{{1, 0}}))
	println(canFinishMy(2, [][]int{{1, 0}, {0, 1}}))

	println(canFinishMy(4, [][]int{{1, 0}, {2, 0}, {3, 2}}))

	println(canFinishMy(1, [][]int{}))
	println(canFinishMy(3, [][]int{{1, 0}, {0, 2}, {2, 1}}))
}

func canFinish(
	numCourses int,
	prerequisites [][]int) bool {

	fmt.Print("courses:")
	fmt.Println(prerequisites)

	G := make([][]int, numCourses)	
	degree := make([]int, numCourses)

	for _, item := range prerequisites {
		i := item[0]
		j := item[1]
		G[i] = append(G[i],j)
		degree[j]++
	}

	fmt.Println("index所依靠的課程們")
	fmt.Print("G:")
	fmt.Println(G)
	fmt.Println("")

	fmt.Println("index被依靠的課程數量")
	fmt.Print("degree:")
	fmt.Println(degree)
	fmt.Println("")
	
	bfs := []int{}

	for index := 0; index < numCourses; index++ {
		if (degree[index] == 0){
			bfs = append(bfs,index)	
		}
	}

	fmt.Println("從來沒被依靠的index")
	fmt.Print("bfs-1:")
	fmt.Println(bfs)
	fmt.Println("")

	for i := 0; i < len(bfs); i++ {
		for _, j := range G[bfs[i]] {
			degree[j]--
			if (degree[j] == 0) {
				bfs = append(bfs, j)
			}
		}
	}

	fmt.Println("處理後所有的課")
	fmt.Print("bfs-2:")
	fmt.Println(bfs)
	fmt.Println("")

	fmt.Println("處理後都沒有index被依靠的課程數量")
	fmt.Print("degree-2:")
	fmt.Println(degree)
	fmt.Print("")

	fmt.Println("len(bfs) == numCourses")
	return len(bfs) == numCourses
}

func Test_canFinish(t *testing.T) {
	//fmt.Println(canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 2}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {2, 0}, {3, 2}, {3, 0}}))

	// println(canFinish(1, [][]int{}))
	// println(canFinish(3, [][]int{{1, 0}, {0, 2}, {2, 1}}))
}