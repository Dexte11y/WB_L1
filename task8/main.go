// Дана переменная int64. Разработать программу которая
// устанавливает i-й бит в 1 или 0.
package main

import (
	"fmt"
)

// Функция для установки i-го бита в 1
// | побитовое или
func setBitToOne(num *int64, i uint) {
	*num |= 1 << i
}

// Функция для установки i-го бита в 0
// & побитовое и
func setBitToZero(num *int64, i uint) {
	*num &= ^(1 << i)
}

func main() {
	var num int64 = 42
	fmt.Printf("Исходное число: %d\n", num)

	// Установить 4-й бит в 1
	setBitToOne(&num, 4)
	fmt.Printf("Число после установки 4-го бита в 1: %d\n", num)

	// Установить 5-й бит в 0
	setBitToZero(&num, 5)
	fmt.Printf("Число после установки 5-го бита в 0: %d\n", num)
}
