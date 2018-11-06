package main

import "fmt"

func main() {
	flowerbed := []int{1,0,0,0,0,1}
	n := 2
	fmt.Println(canPlaceFlowers(flowerbed,n))
}
func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	for i := 0; i < l; i++ {
		if flowerbed[i] == 0 && ((i+1 < l && flowerbed[i+1] == 0) || i+1 >= l) && ((i-1 >= 0 && flowerbed[i-1] == 0) || i-1 < 0) {
			flowerbed[i] = 1
			n--
			if n <= 0 {
				return true
			}
		}
	}

	return n <= 0
}