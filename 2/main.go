//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"os"
)

func SquareElem(array []int, ch chan (int)) {
	for _, v := range array {
		go func(v int) {
			ch <- v * v
		}(v)
	}
}
func main() {
	array := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	go SquareElem(array, ch)
	for range array {
		fmt.Println(os.Stdout, <-ch)
	}
}
