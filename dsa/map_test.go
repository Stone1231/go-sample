package dsa

import (
	"fmt"
	"testing"
)

func Test_Map1(t *testing.T) {
	m := map[int]struct {
		name string
		age  int
	}{
		1: {"user1", 10},
		2: {"user2", 20},
	}
	println(m[1].name)
}
func Test_Map2(t *testing.T) {
	//m := make(map[string]int, 1000)

	m := map[string]int{
		"a": 1,
	}

	if v, ok := m["a"]; ok { //contains a key in go?
		println(ok) //bool
		println(v)
	}
	println(m["c"])
	m["b"] = 2
	delete(m, "c")
	println(len(m))
	for k, v := range m {
		println(k, v)
	}
}
func Test_Map3(t *testing.T) {
	type user struct{ name string }
	m := map[int]user{
		1: {"user1"},
	}
	//m[1].name = "Tom" // Error: cannot assign to m[1].name, map not addressable

	u := m[1]
	u.name = "Tom"
	m[1] = u //    value
	fmt.Println(m[1].name)

	m2 := map[int]*user{
		1: &user{"user1"},
	}
	m2[1].name = "Jack" //指標副本 可以透過指標修改
}

//泆代時可刪除鍵值 但期間有新增操作會有意外
func Test_Map4(t *testing.T) {
	for i := 0; i < 5; i++ {
		m := map[int]string{
			0: "a", 1: "a", 2: "a", 3: "a", 4: "a",
			5: "a", 6: "a", 7: "a", 8: "a", 9: "a",
		}
		for k := range m {
			m[k+k] = "x"
			delete(m, k)
		}
		fmt.Println(m)
	}
}
