package main

import "fmt"

func main() {
	//fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, []int{0, 1}...)
	count := 0
	temp := 0
	// 首补0，尾补0，1，统一为一种情况
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			count = count + (temp-1)/2
			temp = 0
		} else {
			temp++
		}
	}
	return n <= count
}
