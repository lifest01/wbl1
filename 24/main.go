/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// Конструктор
func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// Применяем формулу для расчета расстояния между точками
func GetDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}
func main() {
	point1 := NewPoint(1, 1)
	point2 := NewPoint(4, 4)
	fmt.Printf("%.3f", GetDistance(point1, point2))
}
