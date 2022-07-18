/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

package main

import (
	"fmt"
	"strings"
)

func uniqueSymbol(st string) bool {
	//Приводим строку к нижнему регистру и помещаем в массив рун
	var symbols = []rune(strings.ToLower(st))

	// Проходим циклом по массиву и используя strings.Count проверяем сколько раз символ встречается в строке
	for i := range symbols {
		if strings.Count(string(symbols), string(symbols[i])) > 1 {
			return false
		}
	}

	return true
}

func main() {
	st1 := "abcd"
	st2 := "abCdefAaf"
	st3 := "aabcd"
	fmt.Println(uniqueSymbol(st1))
	fmt.Println(uniqueSymbol(st2))
	fmt.Println(uniqueSymbol(st3))

}
