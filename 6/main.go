package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Закрытие канала
func CloseChannel() chan string {
	ch := make(chan string)
	go func() {

		for {
			select {
			case ch <- "Пример горутины с закрытие канала еще работает":

			case <-ch:
				fmt.Println("Закрытие канала")
				return
			}
		}
	}()
	wg.Done()
	return ch
}

// Используя контекст
func CloseCtx() {
	fmt.Println("Context, start")
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Контекст отменился")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("Горутина через контекст еще работает")
			}
		}
	}(ctx)
	time.Sleep(time.Second)
	cancel()
	wg.Done()
}

func main() {
	wg.Add(1)
	myNumber := CloseChannel()
	fmt.Println(<-myNumber)
	fmt.Println(<-myNumber)
	close(myNumber)

	wg.Add(1)
	CloseCtx()

	wg.Wait()
	log.Println("end")
}
