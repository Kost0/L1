package main

import (
	"fmt"
)

func defineType(a interface{}) {
	// Сравниваем тип переменной a с int, string, bool и chan interface{}
	switch a.(type) {
	case int:
		fmt.Println("Переменная типа int")
	case string:
		fmt.Println("Переменная типа string")
	case bool:
		fmt.Println("Переменная типа bool")
	case chan int:
		fmt.Println("Канал для int")
	case chan string:
		fmt.Println("Канал для string")
	case chan bool:
		fmt.Println("Канал для bool")
	case chan interface{}:
		fmt.Println("Канал для interface")
	default:
		fmt.Println("Тип не определен")
	}
}

func main() {
	a := 1
	b := ""
	c := true
	d := make(chan string)
	defineType(a)
	defineType(b)
	defineType(c)
	defineType(d)
}
