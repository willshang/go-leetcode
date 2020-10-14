package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getKth(10, 20, 5))
}

var m map[int]int

func getKth(lo int, hi int, k int) int {
	m = make(map[int]int)
	arr := make([][2]int, 0)
	for i := lo; i <= hi; i++ {
		arr = append(arr, [2]int{i, getCount(i)})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})
	return arr[k-1][0]
}

func getCount(i int) int {
	res := 0
	temp := i
	for temp != 1 {
		if temp%2 == 1 {
			temp = temp*3 + 1
		} else {
			temp = temp / 2
		}
		res++
		if value, ok := m[temp]; ok {
			res = res + value
			break
		}
	}
	m[i] = res
	return res
}
