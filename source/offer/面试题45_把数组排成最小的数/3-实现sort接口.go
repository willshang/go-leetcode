package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(minNumber([]int{10, 2}))
}

type Arr []string

func (a Arr) Len() int {
	return len(a)
}

func (a Arr) Less(i, j int) bool {
	if a[i]+a[j] < a[j]+a[i] {
		return true
	}
	return false
}

func (a Arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func minNumber(nums []int) string {
	var arr Arr
	for i := 0; i < len(nums); i++ {
		arr = append(arr, strconv.Itoa(nums[i]))
	}
	sort.Sort(arr)
	str := ""
	for i := 0; i < len(arr); i++ {
		str = str + arr[i]
	}
	return str
}
