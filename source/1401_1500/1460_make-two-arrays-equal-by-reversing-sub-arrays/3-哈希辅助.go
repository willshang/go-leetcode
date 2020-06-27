package main

import (
	"fmt"
)

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
}

func canBeEqual(target []int, arr []int) bool {
	m := make(map[int]int, 1001)
	for i := 0; i < len(target); i++ {
		m[target[i]]++
		m[arr[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
