package main

import "fmt"

func main() {
	fmt.Println(canMeasureWater(3, 5, 4))
}

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 || z == x+y {
		return true
	} else if z > x+y || y == 0 {
		return false
	}
	return canMeasureWater(y, x%y, z%y)
}
