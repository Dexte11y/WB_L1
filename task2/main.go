// Написать программу, которая конкурентно рассчитает
// значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	for _, v := range arr {
		go func() {
			squareNumbers := v * v
			fmt.Println(squareNumbers)
		}()
	}
	time.Sleep(2 * time.Second)
}
