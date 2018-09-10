package leetcode

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

	fmt.Println("courses:")
	fmt.Println(prerequisites)

	G := make([][]int, numCourses)//directed graph	
	degree := make([]int, numCourses)//點的邊數

	for _, item := range prerequisites {
		// i -> j
		i := item[0]
		j := item[1]
		G[i] = append(G[i],j)
		degree[j]++
	}

	fmt.Printf("\nG: %v \n", G)

	bfs := []int{}

	for index := 0; index < numCourses; index++ {
		if (degree[index] == 0){
			bfs = append(bfs,index)	
		}
	}

	fmt.Printf("\ndegree: %v \n", degree)
	fmt.Printf("bfs: %v 廣度搜尋無邊資料 degree[index] == 0 \n", bfs)
	
	fmt.Println("\nforLoop bfs:")
	for i := 0; i < len(bfs); i++ {
		fmt.Printf("bfs[%v]: %v \n", i, bfs[i])
		fmt.Printf("	G[%v]: %v \n", bfs[i], G[bfs[i]])
		for _, j := range G[bfs[i]] {
			degree[j]--
			fmt.Printf("		degree[%v]-1: %v \n", j, degree)
			if (degree[j] == 0) {
				//這邊的新增也會進入到bfs for執行
				bfs = append(bfs, j)
				fmt.Printf("			bfs append %v: %v  \n", j, bfs)
			}
		}
		
		fmt.Println("---------")
	}

	fmt.Printf("\nfinally: \n")
	fmt.Printf("degree: %v \n", degree)
	fmt.Printf("bfs: %v \n", bfs)

	check := len(bfs) == numCourses
	fmt.Println("")
	fmt.Println("len(bfs) == numCourses")

	return check
}

func Test_canFinish(t *testing.T) {
	//fmt.Println(canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 2}}))
	
	//fmt.Println(canFinish(5, [][]int{{1, 0}, {2, 0}, {3, 2}, {3, 0}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {2, 0}, {3, 2}, {3, 0}, {4, 1}, {0, 4}}))

	// println(canFinish(1, [][]int{}))
	// println(canFinish(3, [][]int{{1, 0}, {0, 2}, {2, 1}}))
}

func findOrder(
	numCourses int,
	prerequisites [][]int) []int {

	G := make([][]int, numCourses)
	degree := make([]int, numCourses)

	for _, item := range prerequisites {
		i := item[0]
		j := item[1]
		G[i] = append(G[i],j)
		degree[j]++
	}

	bfs := []int{}

	for index := 0; index < numCourses; index++ {
		if (degree[index] == 0){
			bfs = append(bfs,index)	
		}
	}

	//orders := []int{} unsuccess!

	
	// orders := make([]int, len(bfs))
	// copy(orders, bfs)
	orders := append([]int{}, bfs...)//同上

	for i := 0; i < len(bfs); i++ {

		for _, j := range G[bfs[i]] {
			degree[j]--

			if (degree[j] == 0) {
				orders = append([]int{j}, orders...)
				bfs = append(bfs, j)
			}
		}
	}
	check := len(bfs) == numCourses

	if !check{
		return []int{}	
	}

	return orders
}

func Test_findOrder(t *testing.T) {
	//fmt.Println(findOrder(2, [][]int{{1, 0}}))	
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	//fmt.Println(findOrder(5, [][]int{{1, 0}, {2, 0}, {3, 2}, {3, 0}, {4, 1}, {0, 4}}))
}