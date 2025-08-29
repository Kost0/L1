package main

import "fmt"

// Функция для нахождения пересечения, принимает 2 указателя на слайсы
func intersection(nums1 *[]int, nums2 *[]int) *[]int {
	// Слайс для хранения пересечения
	res := make([]int, 0)

	// Map для хранения элементов первого множества
	m := make(map[int]bool)

	// Проходимся по всем элементам первого множества
	for i := 0; i < len(*nums1); i++ {
		// Определяем для ключа равного элементу множества значение true
		m[(*nums1)[i]] = true
	}

	// Проходимся по всем элементам второго множества
	for i := 0; i < len(*nums2); i++ {
		// Если ключ в map равен true
		if m[(*nums2)[i]] {
			// Добавляем его в пересечение
			res = append(res, (*nums2)[i])
		}
	}

	// Возвращаем пересечение
	return &res
}

func main() {
	mas1 := []int{1, 2, 3}
	mas2 := []int{2, 3, 4}

	res := intersection(&mas1, &mas2)
	for _, v := range *res {
		fmt.Println(v)
	}
}
