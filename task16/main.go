// Реализовать быструю сортировку массива (quicksort)
// встроенными методами языка.

package main

import "fmt"

// Функция быстрой сортировки
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	var less, greater []int

	for i, num := range arr {
		if i == pivotIndex {
			continue
		}
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	QuickSort(less)
	QuickSort(greater)

	copy(arr, less)
	arr[len(less)] = pivot
	copy(arr[len(less)+1:], greater)
}

func main() {
	// Пример неотсортированного среза
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Исходный срез:", arr)

	// Сортировка
	QuickSort(arr)

	fmt.Println("Отсортированный срез:", arr)
}
