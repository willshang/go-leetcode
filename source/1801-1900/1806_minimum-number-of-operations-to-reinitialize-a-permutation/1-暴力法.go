package main

import "reflect"

func main() {

}

// leetcode1806_还原排列的最少操作步数
func reinitializePermutation(n int) int {
	res := 0
	target := make([]int, n)
	perm := make([]int, n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	copy(target, perm)
	for {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}
		res++
		if reflect.DeepEqual(target, arr) {
			break
		}
		copy(perm, arr)
	}
	return res
}
