// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	index := 2
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
}
