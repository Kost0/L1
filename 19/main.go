package main

import "fmt"

func main() {
	// Строка превращенная в слайс рун
	str := []rune("hello world😊😀")

	// Проходимся по половине слайса
	for i := 0; i < len(str)/2; i++ {
		// Меняем элементы местами
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}

	// Выводим результат
	fmt.Println(string(str))
}
