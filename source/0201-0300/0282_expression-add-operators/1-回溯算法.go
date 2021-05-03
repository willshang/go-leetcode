package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addOperators("123", 6))
}

// leetcode282_给表达式添加运算符
var res []string

func addOperators(num string, target int) []string {
	res = make([]string, 0)
	dfs(num, target, 0, "", 0, 0)
	return res
}

func dfs(num string, target int, index int, str string, value int, prev int) {
	if index == len(num) {
		if value == target {
			res = append(res, str)
		}
		return
	}
	for i := index; i < len(num); i++ {
		if num[index] == '0' && index < i { // 105 5 => 1*05不符合要求
			return
		}
		s := num[index : i+1]
		a, _ := strconv.Atoi(s)
		if index == 0 {
			dfs(num, target, i+1, str+s, a, a)
		} else {
			dfs(num, target, i+1, str+"+"+s, value+a, a)
			dfs(num, target, i+1, str+"-"+s, value-a, -a)
			dfs(num, target, i+1, str+"*"+s, value-prev+prev*a, prev*a)
		}
	}
}
