package main

import (
	"fmt"
	"index/suffixarray"
	"sort"
)

func main() {
	fmt.Println(multiSearch("mississippi", []string{"is", "ppi", "hi", "sis", "i", "ssippi"}))
}

func multiSearch(big string, smalls []string) [][]int {
	n := len(smalls)
	res := make([][]int, n)
	arr := suffixarray.New([]byte(big)) // 创建后缀树
	for i := 0; i < n; i++ {
		target := []byte(smalls[i])
		temp := arr.Lookup(target, -1) // 返回arr中所有target出现的位置，从后往前
		sort.Ints(temp)
		res[i] = temp
	}
	return res
}
