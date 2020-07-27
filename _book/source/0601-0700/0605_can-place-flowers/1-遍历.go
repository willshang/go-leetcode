package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
}

// leetcode605_种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	// 判断条件
	// 1:当前元素是0
	// 2.前一个元素是0，或者当前是第一个元素
	// 3.后一个元素是0，或者当前是最后一个元素
	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == length-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
			if n <= 0 {
				return true
			}
		}
	}
	return n <= 0
}
