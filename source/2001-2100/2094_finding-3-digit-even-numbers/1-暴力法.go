package main

import "sort"

func main() {

}

// leetcode2094_找出3位偶数
func findEvenNumbers(digits []int) []int {
	m := make(map[int]bool)
	n := len(digits)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i == j || j == k || i == k {
					continue
				}
				v := digits[i]*100 + digits[j]*10 + digits[k]
				if 100 <= v && v%2 == 0 {
					m[v] = true
				}
			}
		}
	}
	arr := make([]int, 0)
	for k := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	return arr
}
