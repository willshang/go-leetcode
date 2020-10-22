package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
}

// leetcode179_最大数
func largestNumber(nums []int) string {
	arr := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		arr = append(arr, strconv.Itoa(nums[i]))
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]+arr[j] >= arr[j]+arr[i]
	})
	res := strings.Join(arr, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}
