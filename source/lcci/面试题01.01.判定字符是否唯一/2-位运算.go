package main

import "fmt"

func main() {
	fmt.Println(isUnique("leetcode"))
}

// 程序员面试金典01.01_判定字符是否唯一
func isUnique(astr string) bool {
	value := uint32(0)
	for i := 0; i < len(astr); i++ {
		index := astr[i] - 'a'
		if value&(1<<index) == (1 << index) {
			return false
		}
		value = value ^ (1 << index)
	}
	return true
}
