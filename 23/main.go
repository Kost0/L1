package main

import (
	"fmt"
	"math/big"
)

// Функция сложения больших чисел
func addBig(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// Функция вычитания больших чисел
func subBig(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// Функция умножения больших чисел
func mulBig(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// Функция деления больших чисел
func divBig(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {
	a := big.NewInt(1000000000000000000)
	b := big.NewInt(100000000000000000)
	fmt.Println(addBig(a, b))
	fmt.Println(subBig(a, b))
	fmt.Println(mulBig(a, b))
	fmt.Println(divBig(a, b))
}
