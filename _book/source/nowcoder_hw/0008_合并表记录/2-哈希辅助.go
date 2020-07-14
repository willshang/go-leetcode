package main

import (
	"fmt"
	"sort"
)

func main() {
	var num int
	for {
		n, _ := fmt.Scanf("%d", &num)
		if n == 0 {
			break
		}
		var key, vale int
		arr := make([]int, 0)
		m := make(map[int]int)
		for i := 0; i < num; i++ {
			fmt.Scanf("%d %d", &key, &vale)
			if _, ok := m[key]; !ok {
				arr = append(arr, key)
			}
			m[key] = m[key] + vale
		}
		sort.Ints(arr)
		for i := 0; i < len(arr); i++ {
			fmt.Printf("%d %d\n", arr[i], m[arr[i]])
		}
	}
}
