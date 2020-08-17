package main

import "fmt"

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}

func subarraysDivByK(A []int, K int) int {
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	res := 0
	for i := 0; i < len(A); i++ {
		sum = sum + A[i]
		value := (sum%K + K) % K
		m[value]++
	}
	for _, v := range m {
		res = res + v*(v-1)/2
	}
	return res
}
