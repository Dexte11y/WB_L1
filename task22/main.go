// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.MaxInt32 // Пример значения, большего чем 2^20
	b := math.MaxInt32 // Пример значения, большего чем 2^20

	// Умножение
	multiplication := a * b
	fmt.Printf("Умножение: %d\n", multiplication)

	// Деление
	division := a / b
	fmt.Printf("Деление: %d\n", division)

	// Сложение
	addition := a + b
	fmt.Printf("Сложение: %d\n", addition)

	// Вычитание
	subtraction := a - b
	fmt.Printf("Вычитание: %d\n", subtraction)
}
