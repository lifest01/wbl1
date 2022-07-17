// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

// Функция проверки наличия значения в слайсе
func contains(sl []string, str string) bool {
	for _, v := range sl {
		// Если значение есть в слайсе значит его нужно добавить в результирующее множество
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	// Объявляем 2 множества
	firstSet := []string{"one", "two", "three", "four", "five"}
	secondSet := []string{"one", "three", "five", "six", "seven"}
	// Объявляем множество которое будет содержать пересечения двух других множеств
	var resultSet []string
	// Проходим по по первому множеству
	for _, v := range firstSet {
		// Если значение есть во втором множестве, значит добавляем это значение в результирующее множество
		if contains(secondSet, v) {
			resultSet = append(resultSet, v)
		}
	}
	fmt.Println(resultSet)
}
