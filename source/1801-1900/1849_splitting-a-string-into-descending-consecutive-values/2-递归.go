package main

import "fmt"

func main() {
	fmt.Println(splitString("543"))
}

// leetcode1849_将字符串拆分为递减的连续值
func splitString(s string) bool {
	return dfs([]byte(s), 0, 0, 0)
}

func dfs(arr []byte, index int, count int, target int) bool {
	if index == len(arr) {
		return count > 1
	}
	value := 0
	for i := index; i < len(arr); i++ {
		value = value*10 + int(arr[i]-'0')
		if count == 0 || value == target-1 {
			if dfs(arr, i+1, count+1, value) == true {
				return true
			}
		}
	}
	return false
}
