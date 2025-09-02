package main

import (
	"fmt"
	"time"
)

// sleep через цикл
func sleepLoop(duration time.Duration) {
	// Время начала блокировки
	start := time.Now()
	// Пока не прошло нужное время цикл продолжает крутиться
	for time.Since(start) < duration {
	}
}

// sleep через канал
func sleepChan(duration time.Duration) {
	// Создаем канал
	done := make(chan struct{})

	// Запускаем горутину
	go func() {
		// Время начала блокировки
		start := time.Now()
		// Пока не прошло нужное время цикл продолжает крутиться
		for time.Since(start) < duration {
		}
		// Закрываем канал
		close(done)
	}()

	// Блокируем в ожидании закрытия канала
	<-done
}

func main() {
	fmt.Println("Stop for 3 seconds")
	sleepLoop(3 * time.Second)
	fmt.Println("Another stop for 3 seconds")
	sleepChan(3 * time.Second)
	fmt.Println("Continue")
}
