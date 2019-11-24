package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(10))
}

/*func mySqrt(x int) int {
	result  := math.Floor(math.Sqrt(float64(x)))
	return int(result)
}*/

/*func mySqrt(x int) int {
	result  := math.Sqrt(float64(x))
	return int(result)
}*/

func mySqrt(x int) int {
	result := x
	for result*result > x {
		fmt.Println(result)
		result = (result + x/result) / 2
	}
	return result
}
