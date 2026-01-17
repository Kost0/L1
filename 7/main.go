package main

import (
	"math/rand"
	"sync"
)

func main() {
	// map для хранения данных
	m := make(map[int]int)

	// Создаем mutex
	mu := &sync.Mutex{}

	// WaitGroup для ожидания горутин
	wg := &sync.WaitGroup{}

	// Увеличиваем счетчик горутин
	wg.Add(1)
	// Запускаем первую горутину
	go func() {
		// Уменьшаем счетчик по окончанию функции
		defer wg.Done()

		for i := 0; i < 10; i++ {
			// Блокируем mutex
			mu.Lock()
			// Безопасно работаем с map
			m[rand.Intn(1000)] = rand.Intn(1000)
			// Открываем mutex
			mu.Unlock()
		}
	}()

	// Увеличиваем счетчик горутин
	wg.Add(1)
	// Запускаем вторую горутину
	go func() {
		// Уменьшаем счетчик по окончанию функции
		defer wg.Done()

		for i := 0; i < 10; i++ {
			// Блокируем mutex
			mu.Lock()
			// Безопасно работаем с map
			m[rand.Intn(1000)] = rand.Intn(1000)
			// Открываем mutex
			mu.Unlock()
		}
	}()

	// Ожидаем выполнения всех горутин
	wg.Wait()
}
