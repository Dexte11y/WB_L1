// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

// Структура счетчика
type Counter struct {
	mu    sync.Mutex     // Мьютекс для безопасного доступа к счетчику
	value int            // Значение счетчика
	wg    sync.WaitGroup // WaitGroup для ожидания завершения всех горутин
}

// Метод для инкрементирования счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	counter := Counter{} // Создаем экземпляр счетчика

	// Создаем несколько горутин для инкрементирования счетчика
	for i := 0; i < 10; i++ {
		counter.wg.Add(1)
		go func() {
			defer counter.wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Ожидаем завершения всех горутин
	counter.wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.value)
}
