package main

import "math/bits"

func main() {

}

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	total := 1 << n
	m := make(map[int]int) // 每门课的先修课
	for i := 0; i < len(relations); i++ {
		a, b := relations[i][0]-1, relations[i][1]-1 // a=>b
		m[b] = m[b] | (1 << a)                       // 第a位置为1
	}
	arr := make([]int, total) // 每种状态的先修课
	for i := 0; i < total; i++ {
		if bits.OnesCount(uint(i)) > k { // 大于k个,标记为-1
			arr[i] = -1
			continue
		}
		for j := i; j > 0; j = (j - 1) & i { // 枚举状态i的子集

		}
	}
}
