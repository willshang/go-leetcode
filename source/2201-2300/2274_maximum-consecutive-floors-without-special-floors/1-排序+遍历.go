package main

import "sort"

func main() {

}

func maxConsecutive(bottom int, top int, special []int) int {
	res := 0
	sort.Ints(special)
	for i := 0; i < len(special); i++ {
		if i == 0 {
			res = max(res, special[i]-bottom) // 跟bottom比
		} else {
			res = max(res, special[i]-special[i-1]-1)
		}
		if i == len(special)-1 { // top跟最大比
			res = max(res, top-special[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
