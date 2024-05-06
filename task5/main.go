// Разработать программу, которая будет последовательно
// отправлять значения в канал, а с другой стороны
// канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	duration := 2 * time.Second
	dataChan := make(chan int)
	defer close(dataChan)

	go func() {
		for i := 0; ; i++ {
			dataChan <- i
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range dataChan {
			fmt.Println(data)
		}
	}()

	<-time.After(duration)
	wg.Wait()
}
