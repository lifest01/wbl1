// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := []int{5, 16, 66, 13, -144, 8, 6, 2, 18, 4}
	fmt.Println(quicksort(array))
}

func quicksort(a []int) []int {
	// если элемент один то сортировать в нем нечего
	if len(a) < 2 {
		return a
	}
	// берем крайние элементы массива
	left, right := 0, len(a)-1
	// Выбираем опорный элемент
	mid := rand.Int() % len(a)
	// Ставим опорный элемент в конец
	a[mid], a[right] = a[right], a[mid]

	for i, _ := range a {
		// Ищем элементы которые меньше опорного
		if a[i] < a[right] {
			// Если нашли то ставим этот элемент в место опороного и сдвигаем индекс возвращения опорного элемента
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	// возвращаем опорный элемент на место
	a[left], a[right] = a[right], a[left]

	// сортируем подмассивы которые меньше и больше опорного
	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
