package main

import "sort"

func main() {

}

func findOriginalArray(changed []int) []int {
	res := make([]int, 0)
	n := len(changed)
	if n%2 == 1 {
		return nil
	}
	sort.Ints(changed)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[changed[i]]++
	}
	for i := 0; i < n; i++ {
		if m[changed[i]] != 0 {
			res = append(res, changed[i])
			m[changed[i]]--
			m[changed[i]*2]--
		}
	}
	if len(res)*2 != n {
		return nil
	}
	return res
}
