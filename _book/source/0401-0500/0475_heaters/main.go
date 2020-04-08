package main

import (
	"fmt"
	"sort"
)

func main() {
	houses := []int{4, 5, 6, 7, 8, 9, 10}
	heaters := []int{2, 4, 6, 8}
	fmt.Println(findRadius(houses, heaters))
}
func findRadius(houses []int, heaters []int) int {
	if len(heaters) == 0 {
		return 0
	}

	sort.Ints(houses)
	sort.Ints(heaters)
	min := 0
	j := 0
	for i := 0; i < len(houses); i++ {
		for j < len(heaters)-1 && Abs(houses[i], heaters[j]) >= Abs(houses[i], heaters[j+1]) {
			fmt.Println(i, j)
			j++
		}
		min = Max(Abs(houses[i], heaters[j]), min)
	}

	return min

}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
