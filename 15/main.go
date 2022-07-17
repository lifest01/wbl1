/* К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

package main

import (
	"fmt"
)

var justString string

func main() {
	someFunc()
}

func someFunc() {
	v := createHugeString(1 << 10) // 1024 символа а
	justString = v[:100]
	// Получили срез byte, так как символы кириллицы занимают два байта то получили в итоге 50 символов.
	fmt.Println(justString)

	// Исправленный вариант
	justString = string([]rune(v)[:100])
	fmt.Println(justString)

}

// реализация createHugeString
func createHugeString(lenStr int) string {
	m := make([]rune, lenStr, lenStr)
	for i := 0; i < lenStr; i++ {
		m[i] = 'п'
	}
	return string(m)
}
