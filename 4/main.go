package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Создаем контекст для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Определяем количество воркеров
	workers := 10

	// Создаем канал для чтения сигналов
	sigChan := make(chan os.Signal, 1)
	// Подписываемся на получение сигнала SIGINT (ctrl + c)
	signal.Notify(sigChan, syscall.SIGINT)

	// Канал для записи int
	ch := make(chan int)

	// Горутина для записи данных в канал
	go func() {
		// Бесконечный цикл для постоянной записи
		for {
			select {
			// случай отмены контекста
			case <-ctx.Done():
				// Закрываем канал
				close(ch)
				// Завершаем горутину
				return
			// Остальные случаи
			default:
				// Записываем случайное число в канал
				ch <- rand.Intn(1000)
				// Для удобства чтения результата делаем небольшую задержку
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// WaitGroup для работы с горутинами
	wg := sync.WaitGroup{}

	// Запускаем нужное количество воркеров
	for i := 0; i < workers; i++ {
		// Увеличиваем счетчик в WaitGroup
		wg.Add(1)
		// Запускаем горутину-воркер
		go func() {
			// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
			defer wg.Done()
			// Читаем данные из канала
			// После закрытия канала завершится и горутина
			for n := range ch {
				// Выводим результат и номер воркера, чтобы показать что данные обрабатывают разные горутины
				fmt.Printf("Воркер номер %d, прочитал: %d\n", i, n)
			}
		}()
	}

	// Получаем сигнал о завершении работы
	<-sigChan
	// Отменяем контекст
	cancel()

	// Ожидаем окончания всех горутин
	wg.Wait()
	// Вывод обозначающий корректное завершение
	fmt.Println("Graceful shutdown")
}
