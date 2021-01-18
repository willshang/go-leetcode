package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(splitIntoFibonacci("123456579"))
}

// leetcode842_将数组拆分成斐波那契序列
var res []int

func splitIntoFibonacci(S string) []int {
	res = make([]int, 0)
	dfs(S, 0, 0, 0, make([]int, 0))
	return res
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
		if value > math.MaxInt32 {
			break
		}
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
