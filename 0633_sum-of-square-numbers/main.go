package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(5))
}

func judgeSquareSum(c int) bool {
	if c < 0{
		return false
	}

	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right{
		current := left*left + right*right
		if current < c{
			left++
		}else if current > c{
			right--
		}else {
			return true
		}
	}
	return false
}

/*func judgeSquareSum(c int) bool {
	a := intSqrt(c)
	for a >= 0{
		if isSquare(c - a*a){
			return true
		}
		a--
	}
	return false
}

func intSqrt(c int)int  {
	return int(math.Sqrt(float64(c)))
}

func isSquare(b2 int)bool  {
	b := intSqrt(b2)
	return b*b == b2
}*/