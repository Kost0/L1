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

	// Запускаем первую горутину
	go func() {
		for i := 0; i < 10; i++ {
			// Блокируем mutex
			mu.Lock()
			// Безопасно работаем с map
			m[rand.Intn(1000)] = rand.Intn(1000)
			// Открываем mutex
			mu.Unlock()
		}
	}()

	// Запускаем вторую горутину
	go func() {
		for i := 0; i < 10; i++ {
			// Блокируем mutex
			mu.Lock()
			// Безопасно работаем с map
			m[rand.Intn(1000)] = rand.Intn(1000)
			// Открываем mutex
			mu.Unlock()
		}
	}()
}
