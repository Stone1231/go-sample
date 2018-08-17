package typeex

import (
	"fmt"
	"testing"
	"unsafe"
)

const x, y int = 1, 2
const s = "Hello, World!"

// 多常量初始化 // 类型推断
// 常量组
const (
	ci1, ci2      = 10, 100
	cb1      bool = false
)

const (
	cs1 = "abc"
	cs2 // x2 = "abc"
)

const (
	cs3 = "abc"
	ci4 = len(cs3)
	cu1 = unsafe.Sizeof(ci2)
)

const (
	cby1 byte = 100 // int to byte
	//ci5 int  = 1e20 // float64 to int, overflows
)

const (
	Sunday    = iota //0
	Monday           // 1,通常省略后续 表达式。
	Tuesday          //2
	Wednesday        //3
	Thursday         //4
	Friday           //5
	Saturday         //6
)

const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) //iota=1
	MB                          //與KB表達相同, iota = 2
	GB
	TB
)

func Test_main(t *testing.T) {
	const x = "xxx" // 未使 局部常量不会引发编译错误。
	fmt.Println(cs2)
	fmt.Println(Saturday)
	fmt.Println(MB)
	fmt.Println(17 & 3)
}

const a = 2 + 3.0        // a == 5.0   (untyped floating-point constant)
const b = 15 / 4         // b == 3     (untyped integer constant)
const c = 15 / 4.0       // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3 / 2  // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3 / 2. // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0       // d == 8     (untyped integer constant)
const e = 1.0 << 3       // e == 8     (untyped integer constant)
// const f = int32(1) << 33  // illegal    (constant 8589934592 overflows int32)
// const g = float64(2) >> 1 // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar"  // h == true  (untyped boolean constant)
const j = true           // j == true  (untyped boolean constant)
const k = 'w' + 1        // k == 'x'   (untyped rune constant)
const l = "hi"           // l == "hi"  (untyped string constant)
const m = string(k)      // m == "x"   (type string)
const Σ = 1 - 0.707i     //            (untyped complex constant)
const Δ = Σ + 2.0e-4     //            (untyped complex constant)
const Φ = iota*1i - 1/1i //            (untyped complex constant)
