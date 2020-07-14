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
		arr := make([]string, 0)
		var str string
		for i := 0; i < num; i++ {
			fmt.Scanf("%s", &str)
			arr = append(arr, str)
		}
		sort.Strings(arr)
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	}
}
