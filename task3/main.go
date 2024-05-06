// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	results := make(chan int)
	done := make(chan struct{})

	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- n * n
		}(num)
	}

	go func() {
		wg.Wait()      // Ждем, пока все горутины завершат свою работу
		close(results) // Закрываем канал results после завершения вычислений
	}()

	go func() {
		defer close(done)
		sum := 0
		for n := range results {
			sum += n
		}
		fmt.Println("Сумма квадратов:", sum)
	}()

	<-done // Ждем завершения вычислений
}
