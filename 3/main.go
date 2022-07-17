//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных
//вычислений.

package main

import "fmt"

func SumSquare(array []int, ch chan int) {
	for _, v := range array {
		go func(v int) {
			ch <- v * v
		}(v)
	}
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	var total int
	ch := make(chan int)
	go SumSquare(array, ch)
	for range array {
		fmt.Println(total)
		total += <-ch
	}
	fmt.Println(total)
}
