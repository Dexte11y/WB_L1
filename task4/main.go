// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные
// данные из канала и выводят в stdout. Необходима возможность выбора
// количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	const numWorkers = 5 // Количество воркеров

	// Создаем канал для передачи данных между главным потоком и воркерами
	dataChan := make(chan int)
	// Закрываем канал для остановки работы воркеров
	defer close(dataChan)

	// Запускаем горутину для записи данных в канал
	go func() {
		for i := 1; ; i++ {
			dataChan <- i // Записываем данные в канал
		}
	}()

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Запускаем несколько воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for data := range dataChan {
				fmt.Printf("Worker %d received data: %d\n", id, data)
			}
		}(i)
	}

	// Ждем завершения всех воркеров
	go func() {
		wg.Wait()
		fmt.Println("All workers have finished.")
	}()

	// Перехватываем сигналы ОС (нажатие Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Ожидаем сигнала о завершении программы
	<-sigChan
}
