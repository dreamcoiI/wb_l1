package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	temperatureGroups := make(map[int][]float64)
	for _, temp := range temperatures {
		roundTemp := int(math.Round(temp/10) * 10)
		temperatureGroups[roundTemp] = append(temperatureGroups[roundTemp], temp)
	}

	for temp, values := range temperatureGroups {
		fmt.Printf("%d°C: %v\n", temp, values)
	}
}
