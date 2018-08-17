package dsa

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	//appendValues_old(values[:0], root)
	values = values[:0]
	appendValues(&values, root)
}

//一值clone, 效能差
// appendValuesOld appends the elements of t to values in order
// and returns the resulting slice.
func appendValuesOld(values []int, t *tree) []int {
	if t != nil {
		//fmt.Printf("%d %p\n", t.value, &values)
		values = appendValuesOld(values, t.left)
		//fmt.Printf("%d %p\n", t.value, &values)
		values = append(values, t.value)
		//fmt.Printf("%d %p\n", t.value, &values)
		values = appendValuesOld(values, t.right)
		//fmt.Printf("%d %p\n", t.value, &values)
	}
	return values
}

func appendValues(values *[]int, t *tree) *[]int {
	if t != nil {
		//fmt.Printf("%d %p\n", t.value, values)
		values = appendValues(values, t.left)
		//fmt.Printf("%d %p\n", t.value, values)
		*values = append(([]int)(*values), t.value)
		//fmt.Printf("%d %p\n", t.value, values)
		values = appendValues(values, t.right)
		//fmt.Printf("%d %p\n", t.value, values)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
