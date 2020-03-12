package lockex

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func Test_add(t *testing.T) {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				// ops++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)
}

// CAS
func Test_compareAndSwap(t *testing.T) {
	var value int32

	//不斷地嘗試原子地更新value的值,直到操作成功為止
	addValue := func(delta int32) {
		//在被操作值被頻繁變更的情況下,CAS操作並不那麼容易成功
		//so 不得不利用for循環以進行多次嘗試
		for {
			v := value
			if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
				//在函數的結果值為true時,退出循環
				break
			}
			//操作失敗的緣由總會是value的舊值已不與v的值相等了.
			//CAS操作雖然不會讓某個Goroutine阻塞在某條語句上,但是仍可能會使流產的執行暫時停一下,不過時間大都極其短暫.
		}
	}

	fmt.Println("======old value=======")
	fmt.Println(value)
	fmt.Println("======CAS value=======")
	addValue(3)
	fmt.Println(value)
}

//先比較變量的值是否等於給定舊值，等於舊值的情況下才賦予新值，最後返回新值是否設置成功
func Test_compareAndSwap2(t *testing.T) {
	var sum uint32 = 100
	var wg sync.WaitGroup
	for i := uint32(0); i < 100; i++ {
		wg.Add(1)
		go func(t uint32) {
			defer wg.Done()
			atomic.CompareAndSwapUint32(&sum, 100, sum+1)
		}(i)
	}
	wg.Wait()
	fmt.Println(sum)
}

func Test_load(t *testing.T) {
	var value int32

	addValue := func(delta int32) {
		for {
			//在進行讀取value的操作的過程中,其他對此值的讀寫操作是可以被同時進行的,那麼這個讀操作很可能會讀取到一個只被修改了一半的數據.
			//因此我們要使用載入
			v := atomic.LoadInt32(&value)
			if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
				break
			}
		}
	}

	fmt.Println("======old value=======")
	fmt.Println(value)
	fmt.Println("======CAS value=======")
	addValue(3)
	fmt.Println(value)
}

//存儲某個值時，任何CPU都不會都該值進行讀或寫操作
//存儲操作總會成功，它不關心舊值是什麼，與CAS不同
func Test_store(t *testing.T) {
	var ops int32 = 1
	atomic.StoreInt32(&ops, 2)
	fmt.Println("ops:", ops)
}

// 與CAS操作不同，原子交換操作不會關心被操作的舊值。
// 它會直接設置新值
// 它會返回被操作值的舊值
// 此類操作比CAS操作的約束更少，同時又比原子載入操作的功能更強
func Test_swap(t *testing.T) {
	var e int32
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := atomic.LoadInt32(&e)
			old := atomic.SwapInt32(&e, (tmp + 1))
			fmt.Println("e old : ", old)
		}()
	}
	wg.Wait()
	fmt.Println("e : ", e)
}
