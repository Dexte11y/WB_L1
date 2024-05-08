// Реализовать бинарный поиск встроенными методами языка

package main

import "fmt"

// Функция для выполнения двоичного поиска
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid // Индекс найденного элемента
		} else if arr[mid] < target {
			low = mid + 1 // Искать в правой половине массива
		} else {
			high = mid - 1 // Искать в левой половине массива
		}
	}

	return -1 // Элемент не найден
}

func main() {
	// Пример отсортированного массива
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	target := 8

	// Выполняем двоичный поиск
	index := BinarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
