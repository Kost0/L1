package main

import "fmt"

func oldReverse() {
	// Исходная строка
	str := []rune("snow dog sun")

	// Левый указатель слева и правый указатель слева
	ll, lr := 0, 0

	// Левый указатель справа и правый указатель справа
	rl, rr := len(str)-1, len(str)-1

	// Пока правый указатель слева меньше левого указателя справа
	for lr < rl {
		// Находим край слова слева
		for lr < len(str) && str[lr] != ' ' {
			lr++
		}

		// Находим край слова справа
		for rl > 0 && str[rl] != ' ' {
			rl--
		}

		// Если правый указатель слева больше или равен левому указателю справа
		if lr >= rl {
			// То они указывают на одно и то же слово, которое переставлять не надо
			break
		}

		// Добавляем в слайс все что было до левого слова, правое слово, то что между нужными словами, левое слово, то что после правого слова
		str = []rune(string(str[0:ll]) + string(str[rl+1:rr+1]) + string(str[lr:rl+1]) + string(str[ll:lr]) + string(str[rr+1:]))

		// Перераспределяем указатели, учитывая что длины слов, которые были перемещены, могут отличаться
		lr = ll + rr - rl + 1
		rl = rr - (lr - ll) - 1
		ll = lr
		rr = rl
	}

	// Выводим результат
	fmt.Println(string(str))
}

func reverseWordOrder(s string) string {
	runes := []rune(s)

	n := len(runes)

	reverse(runes, 0, n-1)

	start := 0
	for i := 0; i <= n; i++ {
		if i == n || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func reverse(runes []rune, l, r int) {
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
}

func main() {
	fmt.Println(reverseWordOrder("snow dog sun"))
}
