package main

import "fmt"

func makeSet(mas *[]string) *[]string {
	// Map для определения наличия элемента в множестве
	m := make(map[string]bool)

	// Итоговое множество
	res := make([]string, 0)

	// Проходимся по всем элементам слайса
	for _, str := range *mas {
		// Если элемента до этого не было
		if !m[str] {
			// Добавляем его в множество
			res = append(res, str)
			// Определяем что элемент уже добавлен в множество
			m[str] = true
		}
	}

	// Возвращаем слайс
	return &res
}

func main() {
	mas := []string{"cat", "cat", "dog", "cat", "tree"}

	res := makeSet(&mas)

	for _, v := range *res {
		fmt.Println(v)
	}
}
