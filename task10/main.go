// Дана последовательность температурных колебаний:
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupTemperatures := make(map[int][]float64)
	for _, temp := range temperatures {
		if temp >= 0 {
			key := int(math.Floor(temp/10) * 10)
			groupTemperatures[key] = append(groupTemperatures[key], temp)
		} else {
			key := int(math.Floor(temp/10)*10 + 10)
			groupTemperatures[key] = append(groupTemperatures[key], temp)
		}
	}

	for key, values := range groupTemperatures {
		fmt.Printf("%d:%.1f", key, values)
	}
}
