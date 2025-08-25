package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

// Выход по условию
func exitWithCondition(wg *sync.WaitGroup) {
	// Увеличиваем счетчик в WaitGroup
	wg.Add(1)

	// Запускаем горутину
	go func() {
		// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
		defer wg.Done()
		for i := range 10 {
			// Как только i будет равняться 5 горутина завершится
			if i == 5 {
				fmt.Println("i == 5, goroutine ended")
				return
			}
			fmt.Println(i)
		}
	}()

}

// Выход по сигналу
func exitWithSignal(wg *sync.WaitGroup) {
	// Увеличиваем счетчик в WaitGroup
	wg.Add(1)

	// Создаем канал для чтения сигналов
	sigChan := make(chan os.Signal, 1)
	// Подписываемся на получение сигнала SIGINT
	signal.Notify(sigChan, syscall.SIGINT)

	// Запускаем горутину
	go func() {
		// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
		defer wg.Done()
		// Бесконечный цикл
		for {
			select {
			// В случае получения сигнала SIGINT горутина завершается
			case <-sigChan:
				fmt.Println("Signal received")
				return
			default:
				fmt.Println("Waiting for a signal...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

// Выход через контекст
func exitWithContext(wg *sync.WaitGroup) {
	// Контекст для выхода
	ctx, cancel := context.WithCancel(context.Background())

	// Увеличиваем счетчик в WaitGroup
	wg.Add(1)
	// Запускаем горутину
	go func() {
		// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
		defer wg.Done()
		// Бесконечный цикл
		for {
			select {
			// При отмене контекста завершаем горутину
			case <-ctx.Done():
				fmt.Println("Context cancelled")
				return
			default:
				fmt.Println("Waiting for a context...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	// Отменяем контекст
	cancel()

}

// Выход через контекст с таймером
func exitWithContextTimeout(wg *sync.WaitGroup) {
	// Контекст с таймером
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Увеличиваем счетчик в WaitGroup
	wg.Add(1)

	// Запускаем горутину
	go func() {
		// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
		defer wg.Done()
		for {
			select {
			// По истечении таймера завершаем горутину
			case <-ctx.Done():
				fmt.Println("Timer cancelled")
				return
			default:
				fmt.Println("Waiting for a timer...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

// Выход с помощью runtime.Goexit()
func exitWithGoexit(wg *sync.WaitGroup) {
	// Увеличиваем счетчик в WaitGroup
	wg.Add(1)

	// Запускаем горутину
	go func() {
		// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
		defer wg.Done()
		fmt.Println("Waiting for a goexit")
		runtime.Goexit()
	}()
	fmt.Println("Got goexit")
}

// Выход после закрытия канала
func exitWithChanClosed(wg *sync.WaitGroup) {
	// Создание канала
	ch := make(chan int)

	// Увеличиваем счетчик в WaitGroup
	wg.Add(1)
	// Запускаем горутину
	go func() {
		// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
		defer wg.Done()
		// Отправляем данные в канал
		for i := range 10 {
			ch <- i
		}
		// Закрываем канал
		close(ch)
	}()

	// Увеличиваем счетчик в WaitGroup
	wg.Add(1)
	// Запускаем горутину
	go func() {
		// По окончанию выполнения функции в горутине уменьшаем счетчик в WaitGroup
		defer wg.Done()
		// Получаем данные из канала
		for n := range ch {
			fmt.Printf("Channel works, %d\n", n)
		}
		// После закрытия канала горутина завершается
		fmt.Println("Channel ended")
	}()

}

func main() {
	// WaitGroup для работы с горутинами
	wg := sync.WaitGroup{}

	exitWithCondition(&wg)

	exitWithSignal(&wg)

	exitWithContext(&wg)

	exitWithContextTimeout(&wg)

	exitWithGoexit(&wg)

	exitWithChanClosed(&wg)

	// Ожидаем окончания всех горутин
	wg.Wait()
}
