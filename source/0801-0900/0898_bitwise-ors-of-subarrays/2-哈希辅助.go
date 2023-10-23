package main

import "fmt"

func main() {
	fmt.Printf("%b %b %b %d", 5, 7, 5|7, 5|7)
}

func subarrayBitwiseORs(arr []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = true
		for j := i - 1; j >= 0; j-- {
			if arr[j]|arr[i] == arr[j] {
				break
			}
			arr[j] = arr[j] | arr[i]
			m[arr[j]] = true
		}
	}
	return len(m)
}
