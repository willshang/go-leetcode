package main

import "sort"

func main() {

}

func minSetSize(arr []int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	temp := make([][2]int, 0)
	for k, v := range m {
		temp = append(temp, [2]int{k, v})
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i][1] > temp[j][1]
	})
	res := 0
	total := 0
	for i := 0; i < len(temp); i++ {
		total = total + temp[i][1]
		res++
		if total*2 >= len(arr) {
			return res
		}
	}
	return res
}
