package main

import "sort"

func main() {

}

// leetcode2124_检查是否所有A都在B之前
func checkString(s string) bool {
	arr := []byte(s)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return string(arr) == s
}
