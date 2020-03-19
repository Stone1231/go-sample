package initex

import (
	"fmt"
	"testing"
)

// 同個Package底下是不可以有重複的變數或者是函式名稱，但init()可以
// 只要 package 內有 init 的 func，在引入 package 時都會被執行

// 當主程式需要單獨讀取 package 內的 init func 而不讀取額外的變數，這時候就要透過 _ 來讀取 package
// 如果沒有加上 _，當編譯的時候就會報錯，原因就是 main 主程式內沒有用到 pq 內任何非 init() 的功能

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func Test_init(t *testing.T) {
	fmt.Println("test")
}
