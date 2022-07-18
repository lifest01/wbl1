// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
//значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := "73458474610203489237872368294597238282592"
	b := "27356769827398472934678236597698236498121"
	var bigA = new(big.Float)
	var bigB = new(big.Float)
	bigA.SetString(a)
	bigB.SetString(b)

	var result = new(big.Float)

	fmt.Println("Сумма:", result.Add(bigA, bigB))
	fmt.Println("Разность:", result.Sub(bigB, bigA))
	fmt.Println("Умножение:", result.Mul(bigB, bigA))
	fmt.Println("Деление:", result.Quo(bigB, bigA))

}
