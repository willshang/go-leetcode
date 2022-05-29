package main

import "strconv"

func main() {

}

// leetcode2269_找到一个数字的K美丽值
func divisorSubstrings(num int, k int) int {
	res := 0
	str := strconv.Itoa(num)
	for i := k; i <= len(str); i++ {
		v, _ := strconv.Atoi(str[i-k : i])
		if v > 0 && num%v == 0 {
			res++
		}
	}
	return res
}
