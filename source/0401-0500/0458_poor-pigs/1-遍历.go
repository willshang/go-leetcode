package main

import "fmt"

func main() {
	fmt.Println(poorPigs(1000, 15, 59))
}

// leetcode458_可怜的小猪
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	count := minutesToTest/minutesToDie + 1 // 最多喝几次水
	res := 0
	target := 1
	// n^res<buckets
	for target < buckets {
		target = target * count
		res++
	}
	return res
}
