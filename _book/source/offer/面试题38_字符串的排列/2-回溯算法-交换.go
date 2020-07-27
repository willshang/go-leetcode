package main

import "fmt"

func main() {
	//fmt.Println(permutation("abc"))
	fmt.Println(permutation("abb"))
}

// 剑指offer_面试题38_字符串的排列
var res []string

func permutation(s string) []string {
	res = make([]string, 0)
	dfs([]byte(s), 0)
	return res
}

func dfs(arr []byte, index int) {
	if len(arr)-1 == index {
		res = append(res, string(arr))
		return
	}
	m := make(map[byte]bool)
	for i := index; i < len(arr); i++ {
		if _, ok := m[arr[i]]; ok {
			continue
		}
		m[arr[i]] = true
		arr[i], arr[index] = arr[index], arr[i]
		dfs(arr, index+1)
		arr[i], arr[index] = arr[index], arr[i]
	}
}
