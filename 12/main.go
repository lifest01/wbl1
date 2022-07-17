// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

// Функция проверки наличия значения в слайсе
func notContains(sl []string, str string) bool {
	// Если слайс пустой то значит нужно добавить элемент в множество
	if len(sl) == 0 {
		return true
	}
	for _, v := range sl {
		// Если строка уже есть в этом слайсе то значит ее добавлять не нужно
		if v == str {
			return false
		}
	}
	return true
}

func main() {
	// Входная последовательность
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	// Объявление множества которое будет содержать результирующие значения
	var resultSet []string
	for _, v := range array {
		// Если значение не содержится в результирующем множестве значит добавляем его
		if notContains(resultSet, v) {
			resultSet = append(resultSet, v)
		}
	}
	fmt.Println(resultSet)
}
