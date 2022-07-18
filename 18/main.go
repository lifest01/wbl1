/* Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика. */

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	counter int64
}
type mutexCounter struct {
	counter int64
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int64
	counterAtomic := new(atomicCounter)
	counterMutex := new(mutexCounter)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				// потокобезопасный счетчик
				atomic.AddInt64(&counterAtomic.counter, 1)

				// счетчик через блокировки Mutex
				mu.Lock()
				counterMutex.counter++
				mu.Unlock()

				// Обычный счетчик
				counter++
			}
		}()
	}
	wg.Wait()

	fmt.Printf("Mutex counter: %d\n", counterMutex.counter)
	fmt.Printf("Atomic counter: %d\n", counterAtomic.counter)
	fmt.Printf("Casual counter: %d\n", counter)
}
