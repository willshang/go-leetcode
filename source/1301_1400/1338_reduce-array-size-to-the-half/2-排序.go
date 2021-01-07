package main

import "sort"

func main() {

}

// leetcode1338_数组大小减半
func minSetSize(arr []int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	temp := make([]int, 0)
	for _, v := range m {
		temp = append(temp, v)
	}
	sort.Ints(temp)
	res := 0
	total := 0
	for i := len(temp) - 1; i >= 0; i-- {
		total = total + temp[i]
		res++
		if total*2 >= len(arr) {
			return res
		}
	}
	return res
}
