package main

import "fmt"

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}

// leetcode1528_重新排列字符串
func restoreString(s string, indices []int) string {
	arr := []byte(s)
	res := make([]byte, len(s))
	for i := 0; i < len(indices); i++ {
		res[indices[i]] = arr[i]
	}
	return string(res)
}
