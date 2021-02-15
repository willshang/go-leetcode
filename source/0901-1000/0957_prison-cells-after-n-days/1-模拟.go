package main

import "fmt"

func main() {
	//fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 15))
	fmt.Println(prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 15))
}

func prisonAfterNDays(cells []int, N int) []int {
	arr := [8]int{}
	for i := 0; i < 8; i++ {
		arr[i] = cells[i]
	}
	m := make(map[[8]int]int)
	temp := make([][8]int, 0)
	count := 1
	m[arr] = count
	temp = append(temp, arr)
	for i := 1; i <= N; i++ {
		count++
		next := [8]int{}
		for j := 1; j < 7; j++ {
			if arr[j-1] == arr[j+1] {
				next[j] = 1
			}
		}
		if m[next] > 0 {
			total := len(m) - m[next]
			index := (N - i) % total
			//fmt.Println(i, m[next], total, index, m[next]+index-1, next)
			//for k := 0; k < len(temp); k++ {
			//	fmt.Println(k, temp[k])
			//}
			return getArr(temp[m[next]+index-1])
		}
		arr = next
		m[next] = count
		temp = append(temp, next)
	}
	return getArr(arr)
}

func getArr(arr [8]int) []int {
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		res = append(res, arr[i])
	}
	return res
}
