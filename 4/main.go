//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
//произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ch <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		// Обработка отмены контекста, когда сигнал ctrl+c считался уменьшается счетчик waitgroup
		case <-ctx.Done():
			wg.Done()
			fmt.Printf("Воркер %v остановлен\n", id)
			return
		// Чтение данных из канала
		case data := <-ch:
			fmt.Println("Воркер:", id, "данные:", data)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	// Создание контекста c отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan int)
	// Канал в который будет считываться нажатие ctrl+c
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, syscall.SIGINT)
	// Если нажато ctrl+c то отменяется контекст
	go func() {
		select {
		case <-signalChannel:
			fmt.Println("\nCTRL+C нажата")
			cancel()
		}
	}()

	var workNumber int
	_, err := fmt.Scan("Введите кол-во воркеров", &workNumber)
	if err != nil {
		panic(err)
	}
	// Добавляем в waitGroup количество воркеров
	wg.Add(workNumber)
	// Запускаются воркеры
	for w := 1; w <= workNumber; w++ {
		go worker(w, ch, ctx, &wg)
	}
	// Записываются данные в канал
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(time.Second * 1)
		}
	}()
	// Ждем пока счетчик waitGroup не будет равен 0
	wg.Wait()
}
