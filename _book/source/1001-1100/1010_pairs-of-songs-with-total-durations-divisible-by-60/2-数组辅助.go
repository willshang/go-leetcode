package main

import "fmt"

func main() {
	//fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
	fmt.Println(numPairsDivisibleBy60([]int{20, 40}))
}

// leetcode1010_总持续时间可被60整除的歌曲
func numPairsDivisibleBy60(time []int) int {
	res := 0
	arr := make([]int, 60)
	for i := range time {
		if time[i]%60 == 0 {
			res = res + arr[0]
		} else {
			res = res + arr[60-time[i]%60]
		}
		arr[time[i]%60]++
	}
	return res
}
