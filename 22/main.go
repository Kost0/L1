package main

import "fmt"

func deleteElement(a []int, i int) []int {
	// Сдвигаем хвост на место удаляемого элемента
	copy(a[i:], a[i+1:])
	// Обнуляем последний элемент (для сборщика мусора)
	a[len(a)-1] = 0
	// Уменьшаем длину на 1
	a = a[:len(a)-1]
	// Возвращаем слайс
	return a
}

func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mas = deleteElement(mas, 3)
	fmt.Println(mas)
}
