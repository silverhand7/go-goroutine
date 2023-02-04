package go_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return "Default Value if pool nil"
		},
	}

	pool.Put("Danaerys Targaryen")
	pool.Put("Viserys Targaryen")
	pool.Put("Rhaegar Targaryen")

	group := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	group.Wait()

	fmt.Println("selesai")
}
