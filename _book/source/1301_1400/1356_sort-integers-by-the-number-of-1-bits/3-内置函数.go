package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if bits.OnesCount32(uint32(arr[i])) == bits.OnesCount32(uint32(arr[j])) {
			return arr[i] < arr[j]
		}
		return bits.OnesCount32(uint32(arr[i])) < bits.OnesCount32(uint32(arr[j]))
	})
	return arr
}
