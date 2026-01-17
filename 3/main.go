package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Функция запускающая воркеров
func startWorkers(workers int, ch chan int) {
	// WaitGroup для работы с горутинами
	wg := sync.WaitGroup{}

	// Запускаем нужное количество воркеров
	for i := 0; i < workers; i++ {
		// Увеличиваем счетчик в WaitGroup
		wg.Add(1)
		// Запускаем горутину-воркер
		go func(i int) {
			// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
			defer wg.Done()
			// Читаем данные из канала
			for n := range ch {
				// Выводим результат и номер воркера, чтобы показать что данные обрабатывают разные горутины
				fmt.Printf("Воркер номер %d, прочитал: %d\n", i, n)
			}
		}(i)
	}

	// Ожидаем окончания всех горутин
	wg.Wait()
}

func main() {
	// Канал для записи int
	ch := make(chan int)

	// Определяем количество воркеров
	workers := 10
	// Запускаем воркеры
	go startWorkers(workers, ch)

	// Бесконечный цикл для постоянной записи
	for {
		// Записываем случайное число в канал
		ch <- rand.Intn(1000)
		// Для удобства чтения результата делаем небольшую задержку
		time.Sleep(500 * time.Millisecond)
	}
}
