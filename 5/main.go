package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Время работы в секундах
	n := 10
	// Контекст с таймером
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	// Создаем канал
	ch := make(chan int)

	// WaitGroup для работы с горутинами
	wg := sync.WaitGroup{}

	// Увеличиваем счетчик в WaitGroup
	wg.Add(1)
	// Горутина для записи данных в канал
	go func() {
		// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
		defer wg.Done()
		// Бесконечный цикл для постоянной записи
		for {
			select {
			// Случай окончания таймера
			case <-ctx.Done():
				// Закрываем канал
				close(ch)
				// Завершаем горутину
				return
			// Остальные случаи
			default:
				// Записываем случайное число в канал
				ch <- rand.Intn(1000)
				// Для удобства чтения результата делаем задержку равную секунде
				time.Sleep(time.Second)
			}
		}
	}()

	// Увеличиваем счетчик в WaitGroup
	wg.Add(1)
	// Горутина для чтения данных
	go func() {
		// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
		defer wg.Done()
		// Читаем данные из канала
		for num := range ch {
			// Вывод данных
			fmt.Println(num)
		}
	}()

	// Ожидаем окончания всех горутин
	wg.Wait()
}
