package main

import "fmt"

func main() {
	// Map для хранения по числу массива температур
	m := make(map[int][]float32)

	// Исходный массив температур
	mas := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Проходимся по массиву
	for _, num := range mas {
		// Округляем до десятков
		a := int(num - float32(int(num)%10))
		// По десятку добавляем в массив внутри map
		m[a] = append(m[a], num)
	}

	// Вывод данных
	for k, v := range m {
		fmt.Println(k)
		fmt.Println("{")
		for _, num := range v {
			fmt.Println(num)
		}
		fmt.Println("}")
	}
}
