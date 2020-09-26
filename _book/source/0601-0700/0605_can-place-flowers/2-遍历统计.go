package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	count := 0
	temp := 1
	// 以0开头，计算情况同以0结束，为向中间情况靠齐，可以特殊处理把temp初始化为1
	// 中间计算可以种花，value = (temp-1)/2
	// 最后结束如果为偶数, value=temp/2
	for i := 0; i < length; i++ {
		if flowerbed[i] == 1 {
			count = count + (temp-1)/2
			temp = 0
		} else {
			temp++
		}
	}
	count = count + temp/2
	return n <= count
}
