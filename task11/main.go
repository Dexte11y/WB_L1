// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func intersection(set1, set2 []int) []int {
	setMap := make(map[int]bool)
	intersection := []int{}

	// Добавляем все элементы из первого множества в map
	for _, v := range set1 {
		setMap[v] = true
	}

	// Если элемент из второго множества уже есть в map,
	// то добавляем его в пересечение
	for _, v := range set2 {
		if setMap[v] {
			intersection = append(intersection, v)
		}
	}

	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := intersection(set1, set2)
	fmt.Println("Пересечение множеств:", result)
}
