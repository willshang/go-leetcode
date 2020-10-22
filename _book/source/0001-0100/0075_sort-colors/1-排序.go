package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)

}

func sortColors(nums []int) {
	sort.Ints(nums)
}
