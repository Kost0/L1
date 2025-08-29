package main

import "fmt"

func swapWithAdding(a, b int) (int, int) {
	// В a хранится сумма двух чисел
	a = a + b
	// b = a + b - b = a
	b = a - b
	// a = a + b - a = b
	a = a - b
	return a, b
}

func swapWithXOR(a, b int) (int, int) {
	// На примере a = 2, b = 3
	// a = 10 XOR 11 = 01
	a = a ^ b
	// b = 01 XOR 10 = 11
	b = a ^ b
	// a = 01 XOR 11 = 10
	a = a ^ b
	return a, b
}

func main() {
	a := 5
	b := 4
	fmt.Println("Изначальные значения a и b:")
	fmt.Println(a, b)
	a, b = swapWithAdding(a, b)
	fmt.Println("Меняем их местами с помощью сложения и вычитания:")
	fmt.Println(a, b)
	fmt.Println("Меняем их местами с помощью исключающего или:")
	a, b = swapWithXOR(a, b)
	fmt.Println(a, b)
}
