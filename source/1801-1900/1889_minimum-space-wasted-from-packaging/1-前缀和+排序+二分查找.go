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
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + packages[i-1]
	}
	res := math.MaxInt32
	for i := 0; i < len(boxes); i++ {
		temp := boxes[i]
		sort.Ints(temp)
		if temp[len(temp)-1] < packages[n-1] { // 最大箱子无法装最大包裹
			continue
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
