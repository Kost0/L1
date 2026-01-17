package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Счетчик с использованием mutex
func mutexCounter() {
	// WaitGroup для работы с горутинами
	wg := &sync.WaitGroup{}

	// Создаем mutex
	mu := &sync.Mutex{}

	// Счетчик
	cnt := 0

	// Цикл для запуска горутин
	for i := 0; i < 10; i++ {
		// Увеличиваем счетчик в WaitGroup
		wg.Add(1)

		// Запускаем горутину
		go func() {
			// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
			defer wg.Done()
			// Цикл для инкрементирования счетчика
			for j := 0; j < 10; j++ {
				// Блокируем mutex
				mu.Lock()
				// Инкрементируем счетчик
				cnt++
				// Открываем mutex
				mu.Unlock()
			}
		}()
	}

	// Ожидаем окончания всех горутин
	wg.Wait()

	// Вывод результата (ожидается 100)
	fmt.Println(cnt)
}

// Счетчик с использованием атомарного сложения
func atomicCounter() {
	// WaitGroup для работы с горутинами
	wg := &sync.WaitGroup{}

	// Счетчик
	cnt := int64(0)

	// Цикл для запуска горутин
	for i := 0; i < 10; i++ {
		// Увеличиваем счетчик в WaitGroup
		wg.Add(1)

		// Запускаем горутину
		go func() {
			// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
			defer wg.Done()
			// Цикл для инкрементирования счетчика
			for j := 0; j < 10; j++ {
				// Атомарно инкрементируем счетчик
				atomic.AddInt64(&cnt, 1)
			}
		}()
	}

	// Ожидаем окончания всех горутин
	wg.Wait()

	// Вывод результата (ожидается 100)
	fmt.Println(cnt)
}

// SyncCounter - структура конкурентного счетчика
type SyncCounter struct {
	// счетчик
	cnt int
	// mutex для безопасного изменения cnt
	mu *sync.Mutex
}

// NewSyncCounter - функция конструктор для создания счетчика
func NewSyncCounter() *SyncCounter {
	return &SyncCounter{
		cnt: 0,
		mu:  &sync.Mutex{},
	}
}

// Inc - функция для безопасного инкремента счетчика
func (s *SyncCounter) Inc() {
	s.mu.Lock()
	s.cnt++
	s.mu.Unlock()
}

// startWorkers - функция для демонстрации счетчика
func startWorkers() {
	// WaitGroup для работы с горутинами
	wg := &sync.WaitGroup{}

	// Создаем счетчик
	counter := NewSyncCounter()

	// Цикл для запуска горутин
	for i := 0; i < 10; i++ {
		// Увеличиваем счетчик в WaitGroup
		wg.Add(1)

		// Запускаем горутину
		go func() {
			// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
			defer wg.Done()
			// Цикл для инкрементирования счетчика
			for j := 0; j < 10; j++ {
				// Безопасное увеличение счетчика
				counter.Inc()
			}
		}()
	}

	// Ожидаем окончания всех горутин
	wg.Wait()

	// Вывод результата (ожидается 100)
	fmt.Println(counter.cnt)
}

func main() {
	mutexCounter()

	atomicCounter()

	startWorkers()
}
