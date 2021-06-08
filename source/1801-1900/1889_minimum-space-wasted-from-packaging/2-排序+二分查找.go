package main

import (
	"math"
	"sort"
)

func main() {

}

var mod = 1000000007

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	n := len(packages)
	res := math.MaxInt32
	for i := 0; i < len(boxes); i++ {
		temp := boxes[i]
		sort.Ints(temp)
		if temp[len(temp)-1] < packages[n-1] { // 最大箱子无法装最大包裹
			continue
		}
		sum := 0
		left := 0
		for j := 0; j < len(temp); j++ {

		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
