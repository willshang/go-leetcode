package main

import "fmt"

func main() {
	fmt.Println(findString([]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ta"))
}

// 程序员面试金典10.05_稀疏数组搜索
func findString(words []string, s string) int {
	left := 0
	right := len(words) - 1
	for left <= right {
		mid := left + (right-left)/2
		index := mid
		word := words[mid]
		if word == "" {
			for index = mid; index <= right; index++ {
				if words[index] != "" {
					word = words[index]
					break
				}
			}
		}
		if word == s {
			return index
		} else if word < s {
			left = index + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
