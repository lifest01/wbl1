// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func main() {
	// т.к. бинарный поиск можно использовать только в отсортированном массиве, создаем отсортированный массив
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(array, 7))
}

func binarySearch(a []int, target int) bool {
	low, high := 0, len(a)-1

	// Берем значение по середине, если элемент который мы ищем меньше значит ищем в левой половине иначе в правой и т.д.
	for low <= high {
		mid := (low + high) / 2
		if a[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	// Если массив состоит из одного элемента или найденный элемент не совпадает с тем что мы ищем возвращаем false
	if low == len(a) || a[low] != target {
		return false
	}

	return true
}
