package main

import "fmt"

func main() {
	//fmt.Println(sumSubarrayMins([]int{71, 55, 82, 55}))
	fmt.Println(sumSubarrayMins([]int{11, 11, 11, 11}))
}

func sumSubarrayMins(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		left, right := i-1, i+1
		for ; left >= 0; left-- {
			if arr[left] < arr[i] {
				break
			}
		}
		for ; right < len(arr); right++ {
			if arr[right] <= arr[i] { // 注意边界
				break
			}
		}
		sum := arr[i] * (i - left) * (right - i)
		res = (res + sum) % 1000000007
	}
	return res % 1000000007
}
