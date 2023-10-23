package main

import "fmt"

func main() {
	fmt.Println(shortestSeq([]int{1, 2, 3}, []int{4}))
}

// 程序员面试金典17.18_最短超串
func shortestSeq(big []int, small []int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(small); i++ {
		m[small[i]]++
	}
	total := len(m)
	j := 0
	for i := 0; i < len(big); i++ {
		m[big[i]]--
		if m[big[i]] == 0 {
			total--
		}
		for total == 0 {
			m[big[j]]++
			if m[big[j]] > 0 {
				total++
				if len(res) == 0 || res[1]-res[0] > i-j {
					res = []int{j, i}
				}
			}
			j++
		}
	}
	return res
}
