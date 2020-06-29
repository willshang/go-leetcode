package main

import "fmt"

func main() {
	var num int
	for {
		n, _ := fmt.Scanf("%d", &num)
		if n == 0 {
			break
		}
		var key, vale int
		arr := make([]int, 1024)
		for i := 0; i < num; i++ {
			fmt.Scanf("%d %d", &key, &vale)
			arr[key] = arr[key] + vale
		}
		for i := 0; i < len(arr); i++ {
			if arr[i] != 0 {
				fmt.Printf("%d %d\n", i, arr[i])
			}
		}
	}
}
