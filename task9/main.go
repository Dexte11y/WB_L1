// Разработать конвейер чисел. Даны два канала:
// в первый пишутся числа (x) из массива, во второй —
// результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dataChanFirst := make(chan int)
	dataChanSecond := make(chan int)

	go func() {
		for _, num := range numbers {
			dataChanFirst <- num
		}
		close(dataChanFirst)
	}()

	go func() {
		for num := range dataChanFirst {
			dataChanSecond <- num * 2
		}
		close(dataChanSecond)
	}()

	for data := range dataChanSecond {
		fmt.Println(data)
	}

}
