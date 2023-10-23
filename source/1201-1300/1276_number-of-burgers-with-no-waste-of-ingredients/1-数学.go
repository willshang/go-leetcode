package main

import "fmt"

func main() {
	fmt.Println(numOfBurgers(16, 7))
	fmt.Println(numOfBurgers(17, 4))
	fmt.Println(numOfBurgers(4, 17))
	fmt.Println(numOfBurgers(0, 0))
	fmt.Println(numOfBurgers(2, 1))
}

// leetcode1276_不浪费原料的汉堡制作方案
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	a, b := tomatoSlices, cheeseSlices
	c := a - 2*b
	if c%2 == 0 && c/2 <= b && 0 <= c {
		return []int{c / 2, b - c/2}
	}
	return nil
}
