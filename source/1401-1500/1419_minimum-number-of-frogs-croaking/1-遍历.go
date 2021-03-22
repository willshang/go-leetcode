package main

import "fmt"

func main() {
	fmt.Println(minNumberOfFrogs("croakcroak"))
	fmt.Println(minNumberOfFrogs("crcoakroak"))
}

// leetcode1419_数青蛙
func minNumberOfFrogs(croakOfFrogs string) int {
	res := 0
	var c, r, o, a, k int
	temp := 0
	for _, char := range croakOfFrogs {
		if char == 'c' {
			c++
			if temp > 0 {
				temp-- // 有空闲的青蛙
			} else {
				res++ // 没有空闲的青蛙
			}
		} else if r < c && char == 'r' {
			r++
		} else if o < r && char == 'o' {
			o++
		} else if a < o && char == 'a' {
			a++
		} else if k < a && char == 'k' {
			k++
			temp++ // 结束有空闲
		} else {
			return -1
		}
	}
	if temp != res { // 避免出现"croakcroa"的情况
		return -1
	}
	return res
}
