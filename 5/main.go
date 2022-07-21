//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func Worker(ch chan int, ctx context.Context) {
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-ctx.Done():
			os.Exit(0)
		}
	}
}

func main() {
	var timeout int
	fmt.Scan("Введите время через которое программа должна завершить работу: ", &timeout)
	// Контекст с таймером по истечении которого программа завершает работу
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()
	ch := make(chan int)
	// Воркер который считывает
	go Worker(ch, ctx)
	// Записываются данные в канал
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
}
