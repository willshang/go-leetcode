package main

import "fmt"

func main() {
	var n int
	for {
		a, _ := fmt.Scan(&n)
		if a == 0 {
			break
		}
		arr := make([][2]int, n)
		for i := 0; i < n; i++ {
			var x, y int
			fmt.Scan(&x, &y)
			arr[i] = [2]int{x, y}
		}
		fmt.Println(count(arr))
	}
}

func count(arr [][2]int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				// 斜率判断
				l1 := (arr[i][0] - arr[j][0]) * (arr[i][1] - arr[k][1])
				l2 := (arr[i][1] - arr[j][1]) * (arr[i][0] - arr[k][0])
				if l1 != l2 {
					res++
				}
			}
		}
	}
	return res
}
