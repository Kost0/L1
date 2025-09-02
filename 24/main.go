package main

import (
	"fmt"
	"math"
)

// Структура точки
type Point struct {
	// x координата
	x float64
	// y координата
	y float64
}

// Конструктор, принимает координаты
func NewPoint(x, y float64) *Point {
	// Возвращает указатель на объект точки с указанными координатами
	return &Point{x, y}
}

// Функция для подсчета расстояния
func (p Point) Distance(other *Point) float64 {
	// Теорема Пифагора
	return math.Sqrt((p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y))
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(3, 4)
	fmt.Println(a.Distance(b))
}
