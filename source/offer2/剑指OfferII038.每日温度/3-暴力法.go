package main

import (
	"fmt"
)

func main() {
	fmt.Println(dailyTemperatures([]int{73, 72, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(temperatures []int) []int {
	j := 0
	for i := 0; i < len(temperatures); i++ {
		for j = i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				temperatures[i] = j - i
				break
			}
		}
		if j == len(temperatures) {
			temperatures[i] = 0
		}
	}
	return temperatures
}
