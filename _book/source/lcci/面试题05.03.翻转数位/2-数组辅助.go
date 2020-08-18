package main

import "fmt"

func main() {
	fmt.Println(reverseBits(1775))
	fmt.Println(reverseBits(7))
}

func reverseBits(num int) int {
	res := 0
	arr := make([]int, 0)
	count := 0
	for num != 0 {
		if num%2 == 1 {
			count++
		} else {
			arr = append(arr, count)
			count = 0
		}
		num = num / 2
	}
	arr = append(arr, count)
	if len(arr) == 1 {
		return arr[0] + 1
	}
	for i := 1; i < len(arr); i++ {
		res = max(res, arr[i]+arr[i-1])
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
