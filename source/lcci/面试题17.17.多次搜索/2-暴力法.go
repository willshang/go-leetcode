package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiSearch("mississippi", []string{"is", "ppi", "hi", "sis", "i", "ssippi"}))
	//fmt.Println(multiSearch("abc", []string{""}))
}

// 程序员面试金典17.17_多次搜索
func multiSearch(big string, smalls []string) [][]int {
	n := len(smalls)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		arr := make([]int, 0)
		if smalls[i] == "" {
			res[i] = arr
			continue
		}
		for j := 0; j+len(smalls[i]) <= len(big); j++ {
			if big[j:j+len(smalls[i])] == smalls[i] {
				arr = append(arr, j)
			}
		}
		res[i] = arr
	}
	return res
}
