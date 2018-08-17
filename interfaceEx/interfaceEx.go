package interfaceex

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer // 接口嵌入。
	Print()
}

//無任何方法,所有的類別都實現了空介面
func PrintAny(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}

type Attr struct {
	s interface {
		String() string
	}
}

type FuncDo func()

func (self FuncDo) Do() { self() }
