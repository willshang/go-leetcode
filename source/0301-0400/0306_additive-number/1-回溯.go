package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAdditiveNumber("121474836472147483648"))
}

// leetcode306_累加数
var res []int

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	res = make([]int, 0)
	dfs(num, 0, 0, 0, make([]int, 0))
	return len(res) >= 3
}

func dfs(s string, index, sum, prev int, path []int) bool {
	if index == len(s) {
		if len(path) >= 3 {
			res = path
		}
		return len(path) >= 3
	}
	value := 0
	for i := index; i < len(s); i++ {
		// 0开头不满足要求(当前i=index的时候，可以为0， 避免错过1+0=1的情况)
		if s[index] == '0' && i > index {
			break
		}
		value = value*10 + int(s[i]-'0')
		if len(path) >= 2 {
			if value < sum {
				continue
			}
			if value > sum {
				break
			}
		}
		if dfs(s, i+1, prev+value, value, append(path, value)) == true {
			return true
		}
	}
	return false
}
