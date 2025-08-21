package main

import "fmt"

// Структура Human
type Human struct {
	Name    string
	Age     int
	IsAdult bool
}

// Простейший метод
func (h *Human) CheckAge() {
	if h.Age >= 18 {
		h.IsAdult = true
	} else {
		h.IsAdult = false
	}
}

// Структура Action
type Action struct {
	// embedded struct Human
	Human
}

func main() {
	a := Action{
		Human{
			Name:    "Human",
			Age:     20,
			IsAdult: false,
		},
	}
	// Вызываем метод, который изначально принадлежит Human
	a.CheckAge()
	// Видим корректность выполнения
	fmt.Println(a.IsAdult)
}
