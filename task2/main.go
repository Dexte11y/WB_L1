// Написать программу, которая конкурентно рассчитает
// значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n
			fmt.Printf("%d squared is %d\n", n, square)
		}(num)
	}

	wg.Wait() // Ждем завершения всех горутин
}
