package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(crystalBalls([]bool{false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}))
}

func crystalBalls(array []bool) int {
	increment := int(math.Sqrt(float64(len(array))))
	currentMax := increment
	currentMin := 0

	for i := 0; i <= len(array); i += increment {
		if array[i] {
			currentMax = i
			break
		}
		currentMin = i
	}

	for i := currentMin; i < currentMax; i++ {
		if array[i] {
			return i
		}
	}
	return -1
}
