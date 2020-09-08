package main

import "fmt"

func main() {
	fmt.Println(numTimesAllBlue([]int{4, 1, 2, 3}))
}

func numTimesAllBlue(light []int) int {
	res := 0
	sum := 0
	for i := 0; i < len(light); i++ {
		sum = sum + light[i]
		if (i+1)*(i+2)/2 == sum {
			res++
		}
	}
	return res
}
