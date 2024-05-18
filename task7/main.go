// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	data := make(map[int]int)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			data[key] = rand.Intn(100)
		}(i)
	}
	wg.Wait()
	fmt.Println(data)
}
