package main

import "fmt"

func main() {
	fmt.Println(canArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5))
}

func canArrange(arr []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		value := ((arr[i] % k) + k) % k
		m[value]++
	}
	for key, value := range m {
		if key == 0 && value%2 != 0 {
			return false
		}
		target := (k - key) % k // 避免key=0，k-0=k的情况
		if m[target] != value {
			return false
		}
	}
	return true
}
