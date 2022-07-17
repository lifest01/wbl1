// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
// из переменной типа interface{}.

package main

import "fmt"

var (
	ch = make(chan int)
	//входной массив интерфейсов с различными типами
	listInterface = []interface{}{1, 0.87, "string", true, ch}
)

func main() {
	// Перебераем массив и выводим типы переменных используя ключ %T
	for _, v := range listInterface {
		fmt.Printf("%T\n", v)
	}
}
