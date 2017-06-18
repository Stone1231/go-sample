package main

func main() {

	//varastruct{xint}={100} //syntaxerror
	// var b []int = { 1, 2, 3 }
	// c := struct {x int; y string}
	// {
	// }
	var a = struct{ x int }{100}
	var b = []int{1, 2, 3}

	c := []int{1,
		2,
	}

	d := []int{1,
		2}
	//, or } 結尾
}
