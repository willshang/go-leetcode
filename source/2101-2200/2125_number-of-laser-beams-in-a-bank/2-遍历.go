package main

import "strings"

func main() {

}

// leetcode2125_银行中的激光束数量
func numberOfBeams(bank []string) int {
	res := 0
	prev := 0
	for i := 0; i < len(bank); i++ {
		cur := strings.Count(bank[i], "1")
		if cur > 0 {
			res = res + cur*prev
			prev = cur
		}
	}
	return res
}
