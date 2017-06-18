package main

func main() {
	test1()
	test2()
}

type bigint int64 //自訂新類型
func test1() {
	var x bigint = 100
	println(x)
}

func test2() {
	x := 1234
	var b bigint = bigint(x)
	var b2 int64 = int64(b)
	//var s myslice = []int{1, 2, 3}  //error
	//var s2 []int = s
}
