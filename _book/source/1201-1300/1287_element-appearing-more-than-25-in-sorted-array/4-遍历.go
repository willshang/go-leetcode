package main

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 7, 10}))
	//fmt.Println(findSpecialInteger([]int{1}))
}

func findSpecialInteger(arr []int) int {
	length := len(arr) / 4
	for i := 1; i <= 2; i++ {
		value := arr[length*i]
		left := length * i
		for left > 0 {
			if arr[left] == arr[left-1] {
				left--
			} else {
				break
			}
		}
		right := length * i
		for right < len(arr)-1 {
			if arr[right] == arr[right+1] {
				right++
			} else {
				break
			}
		}
		if right-left+1 > length {
			return value
		}
	}
	return arr[3*length]
}
