//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
//произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	go func() {
		for i := 0; ; i++ {
			fmt.Println(i)
		}
		defer wg.Done()
	}()

}
