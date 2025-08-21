package main

import (
	"fmt"
	"sync"
)

func main() {
	// Массив из условия
	nums := []int{2, 4, 6, 8, 10}

	// WaitGroup для работы с горутинами
	var wg sync.WaitGroup

	// Проходимся по всем элементам массива
	for _, num := range nums {
		// Увеличиваем счетчик в WaitGroup
		wg.Add(1)
		// Запускаем горутину для каждого элемента массива
		go func() {
			// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
			defer wg.Done()
			// Выводим квадрат
			fmt.Println(num * num)
		}()
	}

	// Ожидаем окончания всех горутин
	wg.Wait()
}
