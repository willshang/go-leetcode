package main

import "sort"

func main() {

}

func maxmiumScore(cards []int, cnt int) int {
	a, b := make([]int, 0), make([]int, 0)
	for i := 0; i < len(cards); i++ {
		if cards[i]%2 == 0 {
			a = append(a, cards[i])
		} else {
			b = append(b, cards[i])
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	x, y := len(a), len(b)
	arrA, arrB := make([]int, x+1), make([]int, y+1)
	for i := 0; i < x; i++ {
		arrA[i+1] = arrA[i] + a[i]
	}
	for i := 0; i < y; i++ {
		arrB[i+1] = arrB[i] + b[i]
	}
	res := 0
	for i := 0; i <= cnt; i++ { // 枚举奇数的个数
		n, m := cnt-i, i
		if n <= x && m <= y && (arrA[n]+arrB[m])%2 == 0 {
			res = max(res, arrA[n]+arrB[m])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
