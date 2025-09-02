package main

import (
	"fmt"
	"strings"
)

func checkString(s string) bool {
	// Переводим все символы в нижний регистр
	s = strings.ToLower(s)

	// Map для хранения данных о наличии символа
	m := make(map[byte]bool)

	// Проходимся по всем символам строки
	for _, v := range s {
		// Если он уже был
		if m[byte(v)] {
			// Возвращаем false
			return false
		} else {
			// Иначе запоминаем, что символ уже был
			m[byte(v)] = true
		}
	}

	// Если ни разу символ не повторился, возвращаем true
	return true
}

func main() {
	fmt.Println(checkString("hello world"))
	fmt.Println(checkString("abB"))
	fmt.Println(checkString("abc"))
}
