// Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	var resultStr string
	str := "snow dog sun"
	// через метод Split преобразовываем строку в массив
	sliceStr := strings.Split(str, " ")

	// Проходим по массиву в обратном направлении и через конкатенацию добавляем элементы в строку
	for i := len(sliceStr) - 1; i >= 0; i-- {
		resultStr += sliceStr[i] + " "
	}
	fmt.Println(resultStr)
}
