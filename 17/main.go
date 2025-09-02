package main

import "fmt"

// Итеративный бинарный поиск
func iterBinarySearch(mas []int, target int) int {
	// Левая и правая границы
	low, high := 0, len(mas)-1

	// Пока левый указатель не зашел за правый
	for low <= high {
		// Определяем серединный элемент
		mid := (low + high) / 2
		// Если серединный элемент равен искомому
		if mas[mid] == target {
			// Возвращаем индекс серединного элемента
			return mid
		}
		// Если серединный элемент меньше искомого
		if mas[mid] < target {
			// Ищем правее середины
			low = mid + 1
		}
		// Если серединный элемент больше искомого
		if mas[mid] > target {
			// Ищем левее середины
			high = mid - 1
		}
	}
	// Если указатели пересеклись, значит искомого элемента нет в массиве
	return -1
}

// Рекурсивный бинарный поиск
func recBinarySearch(mas []int, target int) int {
	// Если длина массива равна 0
	if len(mas) == 0 {
		// То искомого элемента нет
		return -1
	}
	// Левая и правая границы
	low, high := 0, len(mas)-1

	// Определяем серединный элемент
	mid := (low + high) / 2
	// Если серединный элемент равен искомому
	if mas[mid] == target {
		// Возвращаем индекс серединного элемента
		return mid
	}
	// Если серединный элемент меньше искомого
	if mas[mid] < target {
		// Ищем правее середины
		res := recBinarySearch(mas[mid+1:], target)
		// Если элемента нет
		if res == -1 {
			// Возвращаем -1
			return -1
		} else {
			// Иначе к результату прибавляем mid + 1, так как индекс при передаче меньшего массива изменится
			return res + mid + 1
		}
	}
	// Если серединный элемент больше искомого
	if mas[mid] > target {
		// Ищем левее середины
		return recBinarySearch(mas[:mid], target)
	}

	return -1
}

func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}

	fmt.Println("Удачный итеративный поиск:", iterBinarySearch(mas, 9))

	fmt.Println("Неудачный итеративный поиск:", iterBinarySearch(mas, 8))

	fmt.Println("Удачный рекурсивный поиск:", recBinarySearch(mas, 3))

	fmt.Println("Неудачный рекурсивный поиск:", recBinarySearch(mas, 11))
}
