// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var mu sync.Mutex
	m := make(map[int]int)
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = rand.Intn(1000)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(m)

}
