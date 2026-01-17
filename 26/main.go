package main

import (
	"fmt"
	"strings"
)

func checkString(s string) bool {
	// Переводим все символы в нижний регистр
	s = strings.ToLower(s)

	// Map для хранения данных о наличии символа
	m := make(map[rune]struct{})

	// Проходимся по всем символам строки
	for _, v := range s {
		// Если он уже был
		if _, ok := m[v]; ok {
			// Возвращаем false
			return false
		}
		// Иначе запоминаем, что символ уже был
		m[v] = struct{}{}
	}

	// Если ни разу символ не повторился, возвращаем true
	return true
}

func main() {
	fmt.Println(checkString("hello world"))
	fmt.Println(checkString("abB"))
	fmt.Println(checkString("abc"))
	fmt.Println(checkString("ппривет"))
	fmt.Println(checkString("привет"))
}
